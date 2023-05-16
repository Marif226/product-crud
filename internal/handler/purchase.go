package handler

import (
	"net/http"
)

func (h *Handler) CreatePurchase(w http.ResponseWriter, r *http.Request) {
	h.services.CreatePurchase()
	w.Write([]byte("Create Purchase!\n"))
}

func (h *Handler) GetPurchaseById(w http.ResponseWriter, r *http.Request) {
	h.services.GetPurchase()
	w.Write([]byte("Get Purchase By Id!\n"))
}

func (h *Handler) UpdatePurchase(w http.ResponseWriter, r *http.Request) {
	h.services.UpdatePurchase()
	w.Write([]byte("Update Purchase!\n"))
}

func (h *Handler) DeletePurchase(w http.ResponseWriter, r *http.Request) {
	h.services.DeletePurchase()
	w.Write([]byte("Delete Buyer!\n"))
}