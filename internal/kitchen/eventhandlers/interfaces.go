package eventhandlers

import (
	"context"

	"github.com/uucoffeeshop/coffeeshop-application/internal/pkg/event"
)

type KitchenOrderedEventHandler interface {
	Handle(context.Context, event.KitchenOrdered) error
}
