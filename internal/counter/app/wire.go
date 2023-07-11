//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/uucoffeeshop/coffeeshop-application/cmd/counter/config"
	"github.com/uucoffeeshop/coffeeshop-application/internal/counter/app/router"
	"github.com/uucoffeeshop/coffeeshop-application/internal/counter/events/handlers"
	"github.com/uucoffeeshop/coffeeshop-application/internal/counter/infras"
	infrasGRPC "github.com/uucoffeeshop/coffeeshop-application/internal/counter/infras/grpc"
	"github.com/uucoffeeshop/coffeeshop-application/internal/counter/infras/repo"
	ordersUC "github.com/uucoffeeshop/coffeeshop-application/internal/counter/usecases/orders"
	"github.com/uucoffeeshop/coffeeshop-application/pkg/postgres"
	"github.com/uucoffeeshop/coffeeshop-application/pkg/rabbitmq"
	pkgConsumer "github.com/uucoffeeshop/coffeeshop-application/pkg/rabbitmq/consumer"
	pkgPublisher "github.com/uucoffeeshop/coffeeshop-application/pkg/rabbitmq/publisher"
	"google.golang.org/grpc"
)

func InitApp(
	cfg *config.Config,
	dbConnStr postgres.DBConnString,
	rabbitMQConnStr rabbitmq.RabbitMQConnStr,
	grpcServer *grpc.Server,
) (*App, func(), error) {
	panic(wire.Build(
		New,
		dbEngineFunc,
		rabbitMQFunc,
		pkgPublisher.EventPublisherSet,
		pkgConsumer.EventConsumerSet,

		infras.BaristaEventPublisherSet,
		infras.KitchenEventPublisherSet,
		infrasGRPC.ProductGRPCClientSet,
		router.CounterGRPCServerSet,
		repo.RepositorySet,
		ordersUC.UseCaseSet,
		handlers.BaristaOrderUpdatedEventHandlerSet,
		handlers.KitchenOrderUpdatedEventHandlerSet,
	))
}

func dbEngineFunc(url postgres.DBConnString) (postgres.DBEngine, func(), error) {
	db, err := postgres.NewPostgresDB(url)
	if err != nil {
		return nil, nil, err
	}
	return db, func() { db.Close() }, nil
}

func rabbitMQFunc(url rabbitmq.RabbitMQConnStr) (*amqp.Connection, func(), error) {
	conn, err := rabbitmq.NewRabbitMQConn(url)
	if err != nil {
		return nil, nil, err
	}
	return conn, func() { conn.Close() }, nil
}
