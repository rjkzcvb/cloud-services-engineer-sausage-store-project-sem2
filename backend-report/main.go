package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	// Health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	// Reports endpoint
	http.HandleFunc("/reports", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		reports := []map[string]interface{}{
			{"id": "1", "user_id": "user1", "action": "view_product", "created_at": "2024-01-01T00:00:00Z"},
			{"id": "2", "user_id": "user2", "action": "add_to_cart", "created_at": "2024-01-01T00:00:00Z"},
		}
		json.NewEncoder(w).Encode(reports)
	})

	// Create report endpoint
	http.HandleFunc("/report", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"status": "created"})
	})

	log.Printf("Backend-report server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
