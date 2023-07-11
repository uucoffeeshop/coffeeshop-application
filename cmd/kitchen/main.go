package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/uucoffeeshop/coffeeshop-application/cmd/kitchen/config"
	"github.com/uucoffeeshop/coffeeshop-application/internal/kitchen/app"
	"github.com/uucoffeeshop/coffeeshop-application/pkg/logger"
	"github.com/uucoffeeshop/coffeeshop-application/pkg/postgres"
	"github.com/uucoffeeshop/coffeeshop-application/pkg/rabbitmq"
	"go.uber.org/automaxprocs/maxprocs"
	"golang.org/x/exp/slog"

	pkgConsumer "github.com/uucoffeeshop/coffeeshop-application/pkg/rabbitmq/consumer"
	pkgPublisher "github.com/uucoffeeshop/coffeeshop-application/pkg/rabbitmq/publisher"

	_ "github.com/lib/pq"
)

func main() {
	// set GOMAXPROCS
	_, err := maxprocs.Set()
	if err != nil {
		slog.Error("failed set max procs", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error("failed get config", err)
	}

	slog.Info("⚡ init app", "name", cfg.Name, "version", cfg.Version)

	// set up logrus
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logger.ConvertLogLevel(cfg.Log.Level))


	a, cleanup, err := app.InitApp(cfg, postgres.DBConnString(cfg.PG.DsnURL), rabbitmq.RabbitMQConnStr(cfg.RabbitMQ.URL))
	if err != nil {
		slog.Error("failed init app", err)
		cancel()
	}

	a.CounterOrderPub.Configure(
		pkgPublisher.ExchangeName("counter-order-exchange"),
		pkgPublisher.BindingKey("counter-order-routing-key"),
		pkgPublisher.MessageTypeName("kitchen-order-updated"),
	)

	a.Consumer.Configure(
		pkgConsumer.ExchangeName("kitchen-order-exchange"),
		pkgConsumer.QueueName("kitchen-order-queue"),
		pkgConsumer.BindingKey("kitchen-order-routing-key"),
		pkgConsumer.ConsumerTag("kitchen-order-consumer"),
	)

	slog.Info("🌏 start server...", "address", fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port))

	go func() {
		err := a.Consumer.StartConsumer(a.Worker)
		if err != nil {
			slog.Error("failed to start Consumer", err)
			cancel()
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		cleanup()
		slog.Info("signal.Notify", v)
	case done := <-ctx.Done():
		cleanup()
		slog.Info("ctx.Done", done)
	}
}
