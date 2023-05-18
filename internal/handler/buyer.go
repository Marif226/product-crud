package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Marif226/product-crud/internal/model"
	"github.com/Marif226/product-crud/pkg/helpers"
)

func (h *Handler) CreateBuyer(w http.ResponseWriter, r *http.Request) {
	var newBuyer model.Buyer
	err := helpers.BindRequestJSON(r, &newBuyer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	id, err := h.services.CreateBuyer(newBuyer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(id)
}

func (h *Handler) GetBuyerById(w http.ResponseWriter, r *http.Request) {
	h.services.GetBuyerById()
	w.Write([]byte("Get Buyer By Id!\n"))
}

func (h *Handler) UpdateBuyer(w http.ResponseWriter, r *http.Request) {
	h.services.UpdateBuyer()
	w.Write([]byte("Update Buyer!\n"))
}

func (h *Handler) DeleteBuyer(w http.ResponseWriter, r *http.Request) {
	h.services.DeleteBuyer()
	w.Write([]byte("Delete Buyer!\n"))
}