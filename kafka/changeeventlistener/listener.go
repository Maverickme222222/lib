// Package changeeventlistener listens to change events on a kafka queue and feeds these events
// to a handler. The listener will keep track of failures and only commit messages that
// have been successfully processed.
package changeeventlistener

import (
	"context"
	"errors"

	"github.com/rs/zerolog"
	"go.elastic.co/apm/v2"

	"github.com/kappapay/backend/lib/golang/src/kappa/common/models"
	"github.com/kappapay/backend/lib/golang/src/kappa/kafka/consumer"
)

// Listener This listener will read ChangeEvent's from a kafka consumer and feeds these messages to a handler
type Listener struct {
	con     *consumer.KappaConsumer
	handler ChangeEventHandler
	errChan chan<- error
	logger  *zerolog.Logger
}

// ChangeEventHandler - handles a generic change event
// A handler should process an event and return nil of the processing was successful.
// If the processing failed, the handler should return an error. The Listener will only commit
// the offset if the processing was successful.
type ChangeEventHandler interface {
	HandleEvent(ctx context.Context, changeEvent *models.ChangeEvent) error
}

// ChangeEventHandlerFunc can be used with NewChangeHandler to get a struct that implements ChangeEventHandler
type ChangeEventHandlerFunc func(ctx context.Context, changeEvent *models.ChangeEvent) error

type funcEventHandler struct {
	f ChangeEventHandlerFunc
}

func (f *funcEventHandler) HandleEvent(ctx context.Context, changeEvent *models.ChangeEvent) error {
	return f.f(ctx, changeEvent)
}

// NewChangeEventHandler can be used to convert a ChangeEventHandlerFun to a ChangeEventHandler
func NewChangeEventHandler(f ChangeEventHandlerFunc) *funcEventHandler {
	return &funcEventHandler{f: f}
}

// NewListener ...
func NewListener(con *consumer.KappaConsumer, changeHandler ChangeEventHandler, logger *zerolog.Logger, errChan chan error) (*Listener, error) {
	return &Listener{
		con:     con,
		handler: changeHandler,
		errChan: errChan,
		logger:  logger,
	}, nil

}

// Start will begin listening on a kafka topic and processing events.
// This method will return immediately and handle processing in a goroutine
func (l *Listener) Start(ctx context.Context) error {

	go func() {
		for {
			err := l.readAndHandleMsg(ctx)
			if err != nil {
				isFatalError := l.handleErr(err)
				if isFatalError {
					return
				}
			}
		}
	}()

	return nil
}

func (l *Listener) readAndHandleMsg(ctx context.Context) error {
	changeEvent := models.ChangeEvent{}
	err := l.con.Read(ctx, &changeEvent)
	if err != nil {
		return err
	}

	l.logger.Info().Msg("Got new message from Kafka")

	// apm start transaction for elastic logging and attach to context
	tx := apm.DefaultTracer().StartTransaction(l.con.Topic, "request")
	defer tx.End()
	ctx = apm.ContextWithTransaction(ctx, tx)
	// attach metadata to the apm transaction
	tx.Context.SetCustom("changeEvent", &changeEvent)

	err = l.handler.HandleEvent(ctx, &changeEvent)
	if err != nil {
		apm.CaptureError(ctx, err).Send() //nolint
		return err
	}

	// success commit the new offset and continue processing new messages
	err = l.con.Commit(ctx)
	if err != nil {
		apm.CaptureError(ctx, err).Send() //nolint
		return err
	}
	l.logger.Info().Msg("Finished handling kafka message")
	return nil

}

func (l *Listener) handleErr(err error) bool {
	// if error is fatal
	if errors.Is(err, context.Canceled) {
		return true
	}

	// if there is a valid error channel then pass the error along for handling by the caller.
	if l.errChan != nil {
		l.errChan <- err
	}
	return false
}
