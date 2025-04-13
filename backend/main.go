package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
)

type Biblia struct {
	ID     int    `json:"id"`
	Nome   string `json:"nome"`
	Versao string `json:"versao"`
	Idioma string `json:"idioma"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "biblia.db")
	if err != nil {
		log.Fatal(err)
	}

	db.Exec("CREATE TABLE IF NOT EXISTS biblias (id INTEGER PRIMARY KEY AUTOINCREMENT, nome TEXT, versao TEXT, idioma TEXT)")

	r := mux.NewRouter()
	r.HandleFunc("/biblias", getBiblias).Methods("GET")
	r.HandleFunc("/biblias", createBiblia).Methods("POST")
	r.HandleFunc("/biblias/{id}", updateBiblia).Methods("PUT")
	r.HandleFunc("/biblias/{id}", deleteBiblia).Methods("DELETE")

	log.Println("Servidor iniciado em http://localhost:8080")
	handler := cors.AllowAll().Handler(r)
	http.ListenAndServe(":8080", handler)
}

func getBiblias(w http.ResponseWriter, r *http.Request) {
	rows, _ := db.Query("SELECT * FROM biblias")
	var biblias []Biblia
	for rows.Next() {
		var b Biblia
		rows.Scan(&b.ID, &b.Nome, &b.Versao, &b.Idioma)
		biblias = append(biblias, b)
	}
	json.NewEncoder(w).Encode(biblias)
}

func createBiblia(w http.ResponseWriter, r *http.Request) {
	var b Biblia
	json.NewDecoder(r.Body).Decode(&b)
	stmt, _ := db.Prepare("INSERT INTO biblias(nome, versao, idioma) VALUES (?, ?, ?)")
	stmt.Exec(b.Nome, b.Versao, b.Idioma)
	w.WriteHeader(http.StatusCreated)
}

func updateBiblia(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var b Biblia
	json.NewDecoder(r.Body).Decode(&b)
	stmt, _ := db.Prepare("UPDATE biblias SET nome=?, versao=?, idioma=? WHERE id=?")
	stmt.Exec(b.Nome, b.Versao, b.Idioma, id)
}

func deleteBiblia(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	stmt, _ := db.Prepare("DELETE FROM biblias WHERE id=?")
	stmt.Exec(id)
}
