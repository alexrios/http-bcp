package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type App struct {
	Router *mux.Router
}

func (a *App) MakeRoutes() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/export/db/{db}/table/{table}", a.Export).Methods("GET")
	router.HandleFunc("/import/db/{db}/table/{table}", a.Import).Methods("GET")
	a.Router = router
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) Export(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	db := vars["db"]
	table := vars["table"]
	result := Export(db, table)
	var p = HSMResponse{Response: result, When: time.Now().Unix()}
	respondWithJSON(w, http.StatusOK, p)
}

func (a *App) Import(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	db := vars["db"]
	table := vars["table"]
	result := Import(db, table)
	var p = HSMResponse{Response: result, When: time.Now().Unix()}
	respondWithJSON(w, http.StatusOK, p)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
