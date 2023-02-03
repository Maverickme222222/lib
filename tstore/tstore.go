// Package tstore provides a generic wrapper that can be used to interact with storage
// objects using database transactions.
package tstore

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// TStore wraps "Store" objects that are backed by sqlx.ExtContext (a db connection). This interface allows the caller to
// either inject a sqlx.DB object or a sqlx.Tx at runtime.
type TStore[Store any] interface {
	Do() Store
	Begin(ctx context.Context) (*sqlx.Tx, Store, error)
}

// StoreGenerator is a function that takes either a sqlx.DB or a sqlx.Tx and returns a Store backed by that connection
// This method will be called by the TStoreImpl when either Do or Begin is called.
type StoreGenerator[Store any] func(sqlx.ExtContext) Store

// Impl implements TStore interface
type Impl[Store any] struct {
	db             *sqlx.DB
	storeGenerator StoreGenerator[Store]
}

// New instantiates a new transactional store. The tstore will use the sqlx.DB and the StoreGenerator to return a Store
// that is backed either by the same database connection or by a database transaction.
func New[Store any](db *sqlx.DB, storeGenerator StoreGenerator[Store]) (*Impl[Store], error) {
	return &Impl[Store]{
		db:             db,
		storeGenerator: storeGenerator,
	}, nil
}

var _ TStore[any] = &Impl[any]{}

// Begin implements TStore
func (t *Impl[Store]) Begin(ctx context.Context) (*sqlx.Tx, Store, error) {
	tx, err := t.db.BeginTxx(ctx, nil)
	if err != nil {
		var storeZeroValue Store
		return nil, storeZeroValue, err
	}
	return tx, t.storeGenerator(tx), nil
}

// Do implements TStore
func (t *Impl[Store]) Do() Store {
	return t.storeGenerator(t.db)
}
