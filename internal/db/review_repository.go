package db

import (
	"database/sql"
	"time"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725731408-team-79536/zadanie-6105.git/internal/models"
)

func CreateReview(db *sql.DB, review models.Review) error {
	query := `INSERT INTO reviews (id, bid_id, rating, comment, created_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := db.Exec(query, review.ID, review.BidID, review.Rating, review.Comment, time.Now())
	return err
}

func GetReviews(db *sql.DB, bidID string) ([]models.Review, error) {
	rows, err := db.Query(`SELECT id, bid_id, rating, comment, created_at FROM reviews WHERE bid_id = $1`, bidID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next() {
		var review models.Review
		if err := rows.Scan(&review.ID, &review.BidID, &review.Rating, &review.Comment, &review.CreatedAt); err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}
	return reviews, nil
}
