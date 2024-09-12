package handlers

import (
	"encoding/json"
	"net/http"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725731408-team-79536/zadanie-6105.git/internal/db"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725731408-team-79536/zadanie-6105.git/internal/models"
	"github.com/google/uuid"
)

func GetTendersHandler(w http.ResponseWriter, r *http.Request) {
	database, err := db.Connect()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer database.Close()

	var tenders []models.Tender
	rows, err := database.Query("SELECT id, name, description, service_type, status, organization_id, created_at FROM tenders")
	if err != nil {
		http.Error(w, "Failed to retrieve tenders", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var tender models.Tender
		if err := rows.Scan(&tender.ID, &tender.Name, &tender.Description, &tender.ServiceType, &tender.Status, &tender.OrganizationID, &tender.CreatedAt); err != nil {
			http.Error(w, "Failed to parse tenders", http.StatusInternalServerError)
			return
		}
		tenders = append(tenders, tender)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tenders)
}

func CreateTenderHandler(w http.ResponseWriter, r *http.Request) {
	var tender models.Tender
	if err := json.NewDecoder(r.Body).Decode(&tender); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Генерируем уникальный идентификатор для тендера
	tender.ID = uuid.New().String()

	// Сохраняем тендер в базе данных
	database, err := db.Connect()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer database.Close()

	if err := db.CreateTender(database, tender); err != nil {
		http.Error(w, "Failed to create tender", http.StatusInternalServerError)
		return
	}

	// Возвращаем результат
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tender)
}

func EditTenderHandler(w http.ResponseWriter, r *http.Request) {
	var tender models.Tender
	if err := json.NewDecoder(r.Body).Decode(&tender); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	database, err := db.Connect()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer database.Close()

	err = db.UpdateTender(database, tender)
	if err != nil {
		http.Error(w, "Failed to update tender", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tender)
}

func RollbackTenderHandler(w http.ResponseWriter, r *http.Request) {
	tenderID := r.URL.Query().Get("id")
	version := r.URL.Query().Get("version")

	database, err := db.Connect()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer database.Close()

	tender, err := db.RollbackTender(database, tenderID, version)
	if err != nil {
		http.Error(w, "Failed to rollback tender", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tender)
}
