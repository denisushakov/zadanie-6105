package main

import (
	"log"
	"net/http"
	"os"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725731408-team-79536/zadanie-6105.git/pkg/handlers"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func main() {
	serverAddress := os.Getenv("SERVER_ADDRESS")
	if serverAddress == "" {
		serverAddress = "0.0.0.0:8080"
	}

	http.HandleFunc("/api/ping", PingHandler)
	http.HandleFunc("/api/tenders", handlers.GetTendersHandler)
	http.HandleFunc("/api/tenders/new", handlers.CreateTenderHandler)
	http.HandleFunc("/api/tenders/edit", handlers.EditTenderHandler)
	http.HandleFunc("/api/tenders/rollback", handlers.RollbackTenderHandler)

	http.HandleFunc("/api/bids/feedback", handlers.CreateFeedbackHandler)
	http.HandleFunc("/api/bids/reviews", handlers.GetReviewsHandler)

	log.Printf("Server is running at %s", serverAddress)
	if err := http.ListenAndServe(serverAddress, nil); err != nil {
		log.Fatal("Server failed: ", err)
	}
}
