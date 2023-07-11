//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/uucoffeeshop/coffeeshop-application/cmd/product/config"
	"github.com/uucoffeeshop/coffeeshop-application/internal/product/app/router"
	"github.com/uucoffeeshop/coffeeshop-application/internal/product/infras/repo"
	productsUC "github.com/uucoffeeshop/coffeeshop-application/internal/product/usecases/products"
	"google.golang.org/grpc"
)

func InitApp(
	cfg *config.Config,
	grpcServer *grpc.Server,
) (*App, error) {
	panic(wire.Build(
		New,
		router.ProductGRPCServerSet,
		repo.RepositorySet,
		productsUC.UseCaseSet,
	))
}
