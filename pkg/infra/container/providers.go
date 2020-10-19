package container

import (
	"github.com/google/wire"
	"product-vault-go/api/handlers"
	"product-vault-go/api/routes"
	"product-vault-go/pkg/domain/repository"
	"product-vault-go/pkg/infra"
)

var ApplicationSet = wire.NewSet(infra.DbSet, repository.RepositoriesSet, handlers.HandlerSet, routes.RoutersSet)
