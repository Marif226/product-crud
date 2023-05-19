package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Marif226/product-crud/internal/model"
	"github.com/Marif226/product-crud/pkg/helpers"
)

func (h *Handler) CreateBuyer(w http.ResponseWriter, r *http.Request) {
	var newBuyer model.Buyer
	err := helpers.BindRequestJSON(r, &newBuyer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.services.CreateBuyer(newBuyer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(id)
}

func (h *Handler) GetAllBuyers(w http.ResponseWriter, r *http.Request) {
	buyersList, err := h.services.GetAllBuyers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(buyersList)
}

func (h *Handler) GetBuyerById(w http.ResponseWriter, r *http.Request) {
	// parse id from the request url
	query := r.URL.Query()
	id, err := strconv.ParseUint(query.Get("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	buyer, err := h.services.GetBuyerById(int(id))
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(buyer)
}

func (h *Handler) UpdateBuyer(w http.ResponseWriter, r *http.Request) {
	var updatedBuyer model.Buyer
	err := helpers.BindRequestJSON(r, &updatedBuyer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	buyer, err := h.services.UpdateBuyer(updatedBuyer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(buyer)
}

func (h *Handler) DeleteBuyer(w http.ResponseWriter, r *http.Request) {
	// parse id from the request url
	query := r.URL.Query()
	id, err := strconv.ParseUint(query.Get("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.services.DeleteBuyer(int(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Buyer and his purchases deleted!"))
}