package main

import (
	"log"
	"net/http"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725731408-team-79536/zadanie-6105.git/config"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725731408-team-79536/zadanie-6105.git/internal/db"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725731408-team-79536/zadanie-6105.git/pkg/handlers"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	database, err := db.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer database.Close()

	handler := &handlers.Handler{DB: database}

	http.HandleFunc("/api/ping", PingHandler)
	http.HandleFunc("/api/tenders", handler.GetTendersHandler)
	http.HandleFunc("/api/tenders/new", handler.CreateTenderHandler)
	http.HandleFunc("/api/tenders/edit", handler.EditTenderHandler)
	http.HandleFunc("/api/tenders/rollback", handler.RollbackTenderHandler)

	http.HandleFunc("/api/bids/feedback", handler.CreateFeedbackHandler)
	http.HandleFunc("/api/bids/reviews", handler.GetReviewsHandler)

	log.Printf("Server is running at %s", cfg.ServerAddress)
	if err := http.ListenAndServe(cfg.ServerAddress, nil); err != nil {
		log.Fatal("Server failed: ", err)
	}
}
