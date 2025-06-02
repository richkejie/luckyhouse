package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	// "os"
)

type Order struct {
	Item     string  `json:"item"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

func main() {
	// Serve static frontend files from the ./static folder
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// API endpoint to receive POSTed orders
	http.HandleFunc("/order", handleOrder)

	port := "8080"
	fmt.Println("Server running at http://localhost:" + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var order Order
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&order)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Log the received order (or process/store it)
	log.Printf("Received order: %+v\n", order)

	// Respond to client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}
