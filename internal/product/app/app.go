package app

import (
	"github.com/uucoffeeshop/coffeeshop-application/cmd/product/config"
	productUC "github.com/uucoffeeshop/coffeeshop-application/internal/product/usecases/products"
	"github.com/uucoffeeshop/coffeeshop-application/proto/gen"
)

type App struct {
	Cfg               *config.Config
	UC                productUC.UseCase
	ProductGRPCServer gen.ProductServiceServer
}

func New(
	cfg *config.Config,
	uc productUC.UseCase,
	productGRPCServer gen.ProductServiceServer,
) *App {
	return &App{
		Cfg:               cfg,
		UC:                uc,
		ProductGRPCServer: productGRPCServer,
	}
}
