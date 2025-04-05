package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type MonkeyFact struct {
	ID      int    `json:"id"`
	Fact    string `json:"fact"`
	Species string `json:"species"`
}

var pool *pgxpool.Pool

func init() {
	var err error
	// Load .env file if it exists
	_ = godotenv.Load()

	// Connect to database
	pool, err = pgxpool.New(context.Background(), os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		log.Fatal(err)
	}

	// Test the connection
	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}

func getRandomFact(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var fact MonkeyFact
	err := pool.QueryRow(context.Background(),
		"SELECT id, fact, species FROM facts ORDER BY RANDOM() LIMIT 1").Scan(&fact.ID, &fact.Fact, &fact.Species)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fact)
}

func getRandomFactBySpecies(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		http.Error(w, "Invalid URL format. Use /fact/species_name", http.StatusBadRequest)
		return
	}
	species := parts[2]

	var fact MonkeyFact
	err := pool.QueryRow(context.Background(),
		"SELECT id, fact, species FROM facts WHERE LOWER(species) = LOWER($1) ORDER BY RANDOM() LIMIT 1",
		species).Scan(&fact.ID, &fact.Fact, &fact.Species)
	if err != nil {
		if err.Error() == "no rows in result set" {
			http.Error(w, "No facts found for this species", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fact)
}

func main() {
	cmd := exec.Command("python3", "read_facts.py")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Println("Running Python script to populate database...")
	if err := cmd.Run(); err != nil {
		log.Printf("Error running Python script: %v", err)
	}

	http.HandleFunc("/fact", getRandomFact)
	http.HandleFunc("/fact/", getRandomFactBySpecies)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
