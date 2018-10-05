package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
}

func (a *App) MakeRoutes() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/export/db/{db}/schema/{schema}/table/{table}", a.Export).Methods("GET")
	router.HandleFunc("/import/db/{db}/schema/{schema}/table/{table}", a.Import).Methods("GET")
	a.Router = router
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) Export(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	db := vars["db"]
	schema := vars["schema"]
	table := vars["table"]
	result, err := Export(db, table)
	if err != nil {
		go DoCallbackRequest(callBackExportUrl, result)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (a *App) Import(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	db := vars["db"]
	schema := vars["schema"]
	table := vars["table"]
	result, err := Import(db, table)
	if err != nil {
		go DoCallbackRequest(callBackImportUrl, result)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
