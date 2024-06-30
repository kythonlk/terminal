package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Drug struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Result string `json:"result"`
}

func main() {
	db, err := sql.Open("sqlite3", "./drugs.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := createTableIfNotExists(db); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var drug Drug
		if err := json.NewDecoder(r.Body).Decode(&drug); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		stmt, err := db.Prepare("INSERT INTO drugs(name, result) VALUES(?, ?)")
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		_, err = stmt.Exec(drug.Name, drug.Result)
		if err != nil {
			http.Error(w, "Failed to save data", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Data saved successfully")
	})

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		rows, err := db.Query("SELECT id, name, result FROM drugs")
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var drugs []Drug
		for rows.Next() {
			var drug Drug
			if err := rows.Scan(&drug.ID, &drug.Name, &drug.Result); err != nil {
				http.Error(w, "Failed to scan data", http.StatusInternalServerError)
				return
			}
			drugs = append(drugs, drug)
		}

		if err := rows.Err(); err != nil {
			http.Error(w, "Error reading data", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(drugs); err != nil {
			http.Error(w, "Failed to encode data", http.StatusInternalServerError)
		}
	})

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createTableIfNotExists(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS drugs (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        result TEXT
    );`
	_, err := db.Exec(query)
	return err
}
