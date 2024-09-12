package handlers

import (
	"encoding/json"
	"net/http"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725731408-team-79536/zadanie-6105.git/internal/db"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725731408-team-79536/zadanie-6105.git/internal/models"
)

func CreateFeedbackHandler(w http.ResponseWriter, r *http.Request) {
	var review models.Review
	if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	database, err := db.Connect()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer database.Close()

	err = db.CreateReview(database, review)
	if err != nil {
		http.Error(w, "Failed to create review", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(review)
}

func GetReviewsHandler(w http.ResponseWriter, r *http.Request) {
	bidID := r.URL.Query().Get("bidID")

	database, err := db.Connect()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer database.Close()

	reviews, err := db.GetReviews(database, bidID)
	if err != nil {
		http.Error(w, "Failed to fetch reviews", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reviews)
}
