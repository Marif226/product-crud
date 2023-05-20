package handler

import (
	"encoding/json"
	"errors"
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

	if newPurchase.Description == "" || newPurchase.Name == "" || newPurchase.Quantity == 0 || newPurchase.Price == 0 || newPurchase.BuyerID == 0 {
		http.Error(w, errors.New("error: empty fields").Error(), http.StatusBadRequest)
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

func (h *Handler) GetAllPurchases(w http.ResponseWriter, r *http.Request) {
	purchasesList, err := h.services.GetAllPurchases()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(purchasesList)
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
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(purchase)
}

func (h *Handler) UpdatePurchase(w http.ResponseWriter, r *http.Request) {
	var updatedPurchase model.Purchase
	err := helpers.BindRequestJSON(r, &updatedPurchase)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if updatedPurchase.Description == "" || updatedPurchase.Name == "" || updatedPurchase.Quantity == 0 || updatedPurchase.Price == 0 || updatedPurchase.BuyerID == 0 {
		http.Error(w, errors.New("error: empty fields").Error(), http.StatusBadRequest)
		return
	}

	purchaseResponse, err := h.services.UpdatePurchase(updatedPurchase)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(purchaseResponse)
}

func (h *Handler) DeletePurchase(w http.ResponseWriter, r *http.Request) {
	// parse id from the request url
	query := r.URL.Query()
	id, err := strconv.ParseUint(query.Get("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.services.DeletePurchase(int(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Purchase deleted!"))
}