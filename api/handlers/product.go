package handlers

import (
	"encoding/json"
	logger "github.com/sirupsen/logrus"
	"net/http"
	"product-vault-go/pkg/domain"
	"product-vault-go/pkg/domain/repository"
)

type ProductHandler struct {
	Repository *repository.PostgresProductRepository
}

func (handler *ProductHandler) Create() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data *domain.ProductData
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&data); err != nil {
			w.WriteHeader(500)
			return
		}
		logger.WithFields(logger.Fields{
			"name": data.Name,
		}).Info("Creating New Product")

		defer r.Body.Close()

		user,dataErr := handler.Repository.Create(data)

		if dataErr != nil {
			w.WriteHeader(500)
			str := `{"error": "error to create user"}`
			errorMessage, err :=json.Marshal([]byte(str))
			if err != nil{
				logger.Error("Error to persist user")
			}
			_, _ = w.Write(errorMessage)
			return
		}
		userJson, errJson := json.Marshal(user)
		if errJson != nil {
			w.WriteHeader(500)
			str := `{"error": "error to format json"}`
			errorMessage, err :=json.Marshal([]byte(str))
			if err != nil{
				logger.Error("error to format json")
			}
			_, _ = w.Write(errorMessage)
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(userJson)

	})
}

func NewProductHandler(userRepository *repository.PostgresProductRepository) *ProductHandler {
	logger.Info("Creating Product Handler Instance")
	return &ProductHandler{userRepository}
}