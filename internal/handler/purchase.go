package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Marif226/product-crud/internal/model"
	"github.com/Marif226/product-crud/pkg/helpers"
)

func (h *Handler) CreatePurchase(w http.ResponseWriter, r *http.Request) {
	var newPurchase model.Purchase
	err := helpers.BindRequestJSON(r, &newPurchase)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.services.CreatePurchase(newPurchase)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(id)
}

func (h *Handler) GetPurchaseById(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id, err := strconv.ParseUint(query.Get("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	purchase, err := h.services.GetPurchaseById(int(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(purchase)
}

func (h *Handler) UpdatePurchase(w http.ResponseWriter, r *http.Request) {
	h.services.UpdatePurchase()
	w.Write([]byte("Update Purchase!\n"))
}

func (h *Handler) DeletePurchase(w http.ResponseWriter, r *http.Request) {
	h.services.DeletePurchase()
	w.Write([]byte("Delete Buyer!\n"))
}