package infras

import (
	"context"

	"github.com/google/wire"
	"github.com/uucoffeeshop/coffeeshop-application/internal/counter/usecases/orders"
	"github.com/uucoffeeshop/coffeeshop-application/pkg/rabbitmq/publisher"
)

var (
	BaristaEventPublisherSet = wire.NewSet(NewBaristaEventPublisher)
	KitchenEventPublisherSet = wire.NewSet(NewKitchenEventPublisher)
)

type (
	baristaEventPublisher struct {
		pub publisher.EventPublisher
	}
	kitchenEventPublisher struct {
		pub publisher.EventPublisher
	}
)

func NewBaristaEventPublisher(pub publisher.EventPublisher) orders.BaristaEventPublisher {
	return &baristaEventPublisher{
		pub: pub,
	}
}

func (p *baristaEventPublisher) Configure(opts ...publisher.Option) {
	p.pub.Configure(opts...)
}

func (p *baristaEventPublisher) Publish(ctx context.Context, body []byte, contentType string) error {
	return p.pub.Publish(ctx, body, contentType)
}

func NewKitchenEventPublisher(pub publisher.EventPublisher) orders.KitchenEventPublisher {
	return &kitchenEventPublisher{
		pub: pub,
	}
}

func (p *kitchenEventPublisher) Configure(opts ...publisher.Option) {
	p.pub.Configure(opts...)
}

func (p *kitchenEventPublisher) Publish(ctx context.Context, body []byte, contentType string) error {
	return p.pub.Publish(ctx, body, contentType)
}
