package main

import (
	"encoding/json"
	"net/http"
	"time"
	"log"
)

// Response структура для JSON-ответа
type Response struct {
	CurrentTime string `json:"current_time"`
}

func currentTimeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	now := time.Now()
	response := Response{
		CurrentTime: now.Format(time.RFC3339), // Форматирование времени в ISO 8601
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("GET /current-time", currentTimeHandler) // Обработчик для GET /current-time

	// Запуск сервера
	log.Fatalf("Server error: %v", http.ListenAndServe(":8080", nil))
}
