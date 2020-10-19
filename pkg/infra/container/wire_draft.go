package container

import (
	"github.com/google/wire"
	"github.com/gorilla/mux"
	"product-vault-go/api/routes"
)

func InitRouter() (*mux.Router, error) {
	wire.Build(ApplicationSet)
	return nil, nil
}

func InitUserRouter() (*routes.ProductRouter, error) {
	wire.Build(ApplicationSet)
	return nil, nil
}