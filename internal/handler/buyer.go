package handler

import "net/http"

func (h *Handler) CreateBuyer(w http.ResponseWriter, r *http.Request) {
	h.services.CreateBuyer()
	w.Write([]byte("Create Buyer!\n"))
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