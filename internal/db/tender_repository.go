package db

import (
	"database/sql"
	"time"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725731408-team-79536/zadanie-6105.git/internal/models"
)

func CreateTender(db *sql.DB, tender models.Tender) error {
	query := `INSERT INTO tenders (id, name, description, status)
              VALUES ($1, $2, $3, $4)`

	_, err := db.Exec(query, tender.ID, tender.Name, tender.Description, tender.Status)

	return err
}

func GetTenders(db *sql.DB) ([]models.Tender, error) {
	rows, err := db.Query(`SELECT id, name, description, status, created_at, updated_at FROM tenders`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tenders []models.Tender
	for rows.Next() {
		var tender models.Tender
		if err := rows.Scan(&tender.ID, &tender.Name, &tender.Description, &tender.Status, &tender.CreatedAt, &tender.UpdatedAt); err != nil {
			return nil, err
		}
		tenders = append(tenders, tender)
	}
	return tenders, nil
}

func UpdateTender(db *sql.DB, tender models.Tender) error {
	query := `UPDATE tenders SET name = $1, description = $2, status = $3, updated_at = $4 WHERE id = $5`

	_, err := db.Exec(query, tender.Name, tender.Description, tender.Status, time.Now(), tender.ID)
	return err
}

func RollbackTender(db *sql.DB, tenderID, version string) (models.Tender, error) {
	var tender models.Tender

	query := `SELECT id, name, description, status, created_at, updated_at FROM tender_versions WHERE id = $1 AND version = $2`
	err := db.QueryRow(query, tenderID, version).Scan(&tender.ID, &tender.Name, &tender.Description, &tender.Status, &tender.CreatedAt, &tender.UpdatedAt)
	if err != nil {
		return tender, err
	}

	// Обновляем текущий тендер до выбранной версии
	updateQuery := `UPDATE tenders SET name = $1, description = $2, status = $3, updated_at = $4 WHERE id = $5`
	_, err = db.Exec(updateQuery, tender.Name, tender.Description, tender.Status, time.Now(), tender.ID)

	return tender, err
}
