package repository

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	logger "github.com/sirupsen/logrus"
	"product-vault-go/pkg/domain"
)

type PostgresProductRepository struct {
	db *sql.DB
}

func (repo PostgresProductRepository) Create(data *domain.ProductData) (domain.Product, error) {
	userId, err := uuid.NewRandom()
	if err != nil {
		logger.Error("Error to generate UUID")
	}
	user := domain.Product{
		Name:  data.Name,
		Id:    userId.String(),
	}
	_, dbErr := repo.db.Exec("INSERT INTO product ( id,name) VALUES($1, $2)", user.Id, user.Name)
	if dbErr != nil {
		logger.Error("Error to persist product",dbErr)
		return user, errors.New("ERROR_IN_DB")
	}
	logger.Info("Product created successfully!!!")
	return user, nil
}

func NewPostgresProductRepository(db *sql.DB) *PostgresProductRepository {
	logger.Info("Creating ProductRepository Instance (PostgreSQL)")
	return &PostgresProductRepository{db}
}
