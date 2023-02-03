package tstore

// In the example below, we create an Entity and add an Identity to the
// newly created Entity. These steps are done inside a db transaction so that
// if adding the identity fails, the entire update is rolled back ( in other
// words the creation of the entity is also rolled back).
// There's also an example of using `Do` which uses the store without a db
// transaction.
// ```
//  import (
// 	"context"

// 	"github.com/jmoiron/sqlx"

// 	"github.com/kappapay/backend/lib/golang/src/kappa/common/models"
// )

// type EntityStore interface {
// 	CreateEntity(ctx context.Context, entity *models.Entity) (*models.Entity, error)
// 	AddIdentity(ctx contenxt.Context, entityID string, identity *models.Identity) error
// 	GetEntityById(ctx context.Context, entityID string) (*models.Entity, error)
// }

// // EntityStoreImpl implements EntityStore (not shown in this example)
// type EntityStoreImpl struct {
// 	db sqlx.ExecerContext
// }

// func NewEntityStore(db sqlx.ExecerContext) EntityStore {
// 	return &EntityStoreImpl{
// 		db: db,
// 	}
// }

// func run() error {
// 	// db is left nil for this example. Normally this would need to be initialized first.
// 	var db *sqlx.DB
// 	entityTStore, err := New(db, NewEntityStore)
// 	if err != nil {
// 		return err
// 	}

// 	ctx := context.Background()

// 	tx, store, err := entityTStore.Begin(ctx)
// 	if err != nil {
// 		return err
// 	}
// 	// This will be a no-op if we commit before returning from this function
// 	defer tx.Rollback()

// 	createdEntity, err := store.CreateEntity(ctx, &models.Entity{})
// 	if err != nil {
// 		return err
// 	}

// 	err = store.AddIdentity(ctx, createdEntity.GetId(), &models.Identity{})
// 	if err != nil {
// This will rollback the entity creation.
// 		return err
// 	}
// 	err = tx.Commit()
// 	if err != nil {
// 		return err
// 	}

// 	entityTStore.Do().GetEntityById(ctx, createdEntity.GetId())
// }
//```
