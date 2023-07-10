package eventhandlers

import (
	"context"

	"github.com/uucoffeeshop/coffeeshop-application/internal/pkg/event"
)

type BaristaOrderedEventHandler interface {
	Handle(context.Context, event.BaristaOrdered) error
}
