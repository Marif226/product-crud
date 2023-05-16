package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Marif226/product-crud/internal/handler"
	"github.com/Marif226/product-crud/internal/repository"
	"github.com/Marif226/product-crud/internal/service"
)

func main() {
	router := http.NewServeMux()
	initRoutes(router)

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

	err := srv.ListenAndServe()
	if err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}

func initRoutes(router *http.ServeMux) {
	r := repository.New()
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