package handler

import (
	"net/http"

	"github.com/Marif226/product-crud/internal/service"
)

type Handler struct {
	services *service.Service
}

func New(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {

}