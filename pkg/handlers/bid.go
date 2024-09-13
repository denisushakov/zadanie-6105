package handlers

import (
	"encoding/json"
	"net/http"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725731408-team-79536/zadanie-6105.git/internal/db"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725731408-team-79536/zadanie-6105.git/internal/models"
)

func (h *Handler) CreateFeedbackHandler(w http.ResponseWriter, r *http.Request) {
	database := h.DB
	var review models.Review
	if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := db.CreateReview(database, review)
	if err != nil {
		http.Error(w, "Failed to create review", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(review)
}

func (h *Handler) GetReviewsHandler(w http.ResponseWriter, r *http.Request) {
	database := h.DB
	bidID := r.URL.Query().Get("bidID")

	reviews, err := db.GetReviews(database, bidID)
	if err != nil {
		http.Error(w, "Failed to fetch reviews", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reviews)
}
