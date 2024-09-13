package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725731408-team-79536/zadanie-6105.git/internal/db"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725731408-team-79536/zadanie-6105.git/internal/models"
	"github.com/google/uuid"
)

func (h *Handler) GetTendersHandler(w http.ResponseWriter, r *http.Request) {
	database := h.DB

	var tenders []models.Tender

	tenders, err := db.GetTenders(database)
	if err != nil {
		log.Printf("Failed to scan tender: %v", err)
		http.Error(w, "Failed to parse tenders", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tenders); err != nil {
		log.Printf("Failed to encode tenders to JSON: %v", err)
		http.Error(w, "Failed to encode tenders", http.StatusInternalServerError)
	}
}

func (h *Handler) CreateTenderHandler(w http.ResponseWriter, r *http.Request) {
	database := h.DB
	var tender models.Tender
	if err := json.NewDecoder(r.Body).Decode(&tender); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	tender.ID = uuid.New().String()

	if err := db.CreateTender(database, tender); err != nil {
		http.Error(w, "Failed to create tender", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tender)
}

func (h *Handler) EditTenderHandler(w http.ResponseWriter, r *http.Request) {
	database := h.DB
	var tender models.Tender
	if err := json.NewDecoder(r.Body).Decode(&tender); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := db.UpdateTender(database, tender)
	if err != nil {
		http.Error(w, "Failed to update tender", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tender)
}

func (h *Handler) RollbackTenderHandler(w http.ResponseWriter, r *http.Request) {
	database := h.DB
	tenderID := r.URL.Query().Get("id")
	version := r.URL.Query().Get("version")

	tender, err := db.RollbackTender(database, tenderID, version)
	if err != nil {
		http.Error(w, "Failed to rollback tender", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tender)
}
