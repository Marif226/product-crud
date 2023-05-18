package main

import (
	"context"
	"database/sql"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Marif226/product-crud/internal/handler"
	"github.com/Marif226/product-crud/internal/repository"
	"github.com/Marif226/product-crud/internal/service"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

func main() {
	// initialize database config
	dbConf, err := initConfig()
	if err != nil {
		log.Fatalf("error during initilizaing the config file, %s", err.Error())
	}

	// Load .env file
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("error during loading .env variable, %s", err.Error())
	}

	// create configured sql.DB instance
	db, err := repository.NewPostgresDB(repository.PGConfig{
		Host: dbConf.Host,
		Port: dbConf.Port,
		Username: dbConf.Username,
		Password: os.Getenv("DB_PASSWORD"),
		DBName: dbConf.DBName,
		SSLMode: dbConf.SSLMode,
	})

	if err != nil {
		log.Fatalf("error during connecting to the database: %s", err.Error())
	}

	// create a router
	router := http.NewServeMux()
	// initialize routes for given router
	initRoutes(router, db)

	srv := &http.Server{
		Addr:           ":8090",
		Handler:        router,
		ReadTimeout:    10 * time.Second, // лимит на чтение запроса в 10 сек
		WriteTimeout:   10 * time.Second, // лимит на запись ответа в 10 сек
		MaxHeaderBytes: 1 << 20, // лимит по памяти на заголовок запроса 
	}

	log.Println("Starting http server...")

	// "Изящное завершение" сервера согласно стандартной документации
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	err = srv.ListenAndServe()
	if err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}

func initRoutes(router *http.ServeMux, db *sql.DB) {
	r := repository.New(db)
	s := service.New(r)
	h := handler.New(s)

	router.HandleFunc("/buyer/create", h.CreateBuyer)
	router.HandleFunc("/buyer/get", h.GetBuyerById)
	router.HandleFunc("/buyer/update", h.UpdateBuyer)
	router.HandleFunc("/buyer/delete", h.DeleteBuyer)

	router.HandleFunc("/purchase/create", h.CreatePurchase)
	router.HandleFunc("/purchase/get", h.GetPurchaseById)
	router.HandleFunc("/purchase/update", h.UpdatePurchase)
	router.HandleFunc("/purchase/delete", h.DeletePurchase)
}

// initialize config file, return error if failed
func initConfig() (*repository.PGConfig, error) {
	var dbConf *repository.PGConfig
	yamlFile, err := os.Open("configs/config.yml")
    if err != nil {
		return dbConf, err
    }
	yamlData, err := io.ReadAll(yamlFile)
	if err != nil {
		return dbConf, err
    }
	
    err = yaml.Unmarshal(yamlData, &dbConf)
    if err != nil {
		return dbConf, err
    }

    return dbConf, nil
}