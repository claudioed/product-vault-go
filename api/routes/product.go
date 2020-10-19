package routes

import (
	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
	"product-vault-go/api/handlers"
)

type ProductRouter struct {
	ProductHandler *handlers.ProductHandler
}

func (userRouter *ProductRouter) CreateProductRoutes(router *mux.Router) {
	router.Handle("/v1/api/products", userRouter.ProductHandler.Create()).Methods("POST")
}

func NewProductRouter(productHandler *handlers.ProductHandler) *ProductRouter {
	logger.Info("Creating Product Router Instance")
	return &ProductRouter{productHandler}
}
