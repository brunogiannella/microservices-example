package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// The person Type (more like an object)
type Pessoa struct {
	Cpf       string `json:"cpf,omitempty"`
	Nome      string `json:"nome,omitempty"`
	Sobrenome string `json:"sobrenome,omitempty"`
	Idade     int    `json:"idade,omitempty"`
}

type Result struct {
	Data Pessoa `json:"data"`
}

// Display a single data
func GetPessoa(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if params["cpf"] == "123" {
		pessoa := Pessoa{
			Cpf:       params["cpf"],
			Nome:      "Juliano",
			Sobrenome: "Silveira",
			Idade:     25}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(Result{Data: pessoa})
		return
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(""))
}

// our main function
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/pessoas/{cpf}", GetPessoa).Methods("GET")
	fmt.Fprintf(os.Stderr, "Servidor iniciado em localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
