package repository

import "product-vault-go/pkg/domain"

type ProductRepository interface {

	Create(data *domain.ProductData) (domain.Product,error)

}
