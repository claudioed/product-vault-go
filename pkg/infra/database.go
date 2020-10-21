package infra

import (
	"database/sql"
	"encoding/json"
	"fmt"
	logger "github.com/sirupsen/logrus"
	"os"
)

type DbCred struct {
	Username string `json:"DB_USER"`
	Password string `json:"DB_PASS"`
}

func NewDatabaseConnection() *sql.DB {
	var name = os.Getenv("DB_NAME")
	var host = os.Getenv("DB_HOST")

	f, err := os.Open("vault/secrets/product-go-vault-sidecar")
	if err != nil {
		logger.Error("error to load db cred %v", err.Error())
		panic(err)
	}

	var cfg DbCred
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		logger.Errorf("fail to parse db cred %v", err.Error())
		panic(err)
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		logger.Errorf("fail to write db cred %v", err.Error())
		panic(err)
	}

	logger.Info("JSON")
	logger.Info(string(data))

	logger.Infof("User %v", cfg.Username)
	dbConnectionString := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", cfg.Username, cfg.Password, host, name)
	logger.Infof("Connection data %s", dbConnectionString)
	db, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
		logger.Error(err)
	}
	return db
}
