package main

import (
	"context"
	"log"
	"net/http"

	"github.com/sankethreddy-404/distributed-order-payment-platform/services/order-service/internal/config"
	"github.com/sankethreddy-404/distributed-order-payment-platform/services/order-service/internal/database"
)

func main() {
	cfg := config.Load()

	ctx := context.Background()

	dbPool, err := database.NewPostgresPool(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer dbPool.Close()

	log.Println("Connected to PostgreSQL successfully")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Order Service is running"))
	})

	log.Println("Order Service starting on port", cfg.Port)

	if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
		log.Fatal(err)
	}
}
