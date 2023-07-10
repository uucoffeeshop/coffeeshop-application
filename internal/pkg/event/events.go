package event

import (
	"time"

	"github.com/google/uuid"

	shared "github.com/uucoffeeshop/coffeeshop-application/internal/pkg/shared_kernel"
)

type BaristaOrdered struct {
	shared.DomainEvent
	OrderID    uuid.UUID       `json:"orderId"`
	ItemLineID uuid.UUID       `json:"itemLineId"`
	ItemType   shared.ItemType `json:"itemType"`
}

func (e BaristaOrdered) Identity() string {
	return "BaristaOrdered"
}

type BaristaOrderUpdated struct {
	shared.DomainEvent
	OrderID    uuid.UUID       `json:"orderId"`
	ItemLineID uuid.UUID       `json:"itemLineId"`
	Name       string          `json:"name"`
	ItemType   shared.ItemType `json:"itemType"`
	TimeIn     time.Time       `json:"timeIn"`
	MadeBy     string          `json:"madeBy"`
	TimeUp     time.Time       `json:"timeUp"`
}

func (e *BaristaOrderUpdated) Identity() string {
	return "BaristaOrderUpdated"
}
