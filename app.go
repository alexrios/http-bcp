package main

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"github.com/pytimer/mux-logrus"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type App struct {
	Router *mux.Router
}

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	govalidator.SetFieldsRequiredByDefault(true)
}

func (a *App) MakeRoutes() {
	routeRegLogger := log.WithFields(log.Fields{
		"phase": "boot",
		"event": "registering route",
	})

	router := mux.NewRouter().StrictSlash(true)
	exportPath := "/export/db/{db}/schema/{schema}/table/{table}"
	router.HandleFunc(exportPath, a.Export).Methods("POST")
	routeRegLogger.WithFields(log.Fields{
		"http-method":      "POST",
		"function-handler": "app.Export",
		"path":             exportPath,
	}).Info("EXPORT")
	importPath := "/import/db/{db}/schema/{schema}/table/{table}"
	router.HandleFunc(importPath, a.Import).Methods("POST")
	routeRegLogger.WithFields(log.Fields{
		"http-method":      "POST",
		"function-handler": "app.Import",
		"path":             importPath,
	}).Info("IMPORT")
	a.Router = router
}

func (a *App) Run(addr string) {
	bootPhaseLogger := log.WithFields(log.Fields{
		"phase": "boot",
	})
	bootPhaseLogger.WithFields(log.Fields{
		"addr": addr,
	}).Info("App started and listening")
	a.Router.Use(muxlogrus.NewLogger().Middleware)
	bootPhaseLogger.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) Export(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	exportReq := ExportRequest{db: vars["db"], schema: vars["schema"], table: vars["table"]}
	_, err := govalidator.ValidateStruct(exportReq)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := Export(exportReq.db, exportReq.schema, exportReq.table)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		go DoCallbackRequest(fmt.Sprintf(callBackExportUrl, exportReq.db, exportReq.schema, exportReq.table), result)
		w.WriteHeader(http.StatusOK)

	}
}

func (a *App) Import(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	importReq := ImportRequest{db: vars["db"], schema: vars["schema"], table: vars["table"]}
	_, err := govalidator.ValidateStruct(importReq)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := Import(importReq.db, importReq.schema, importReq.table)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		go DoCallbackRequest(fmt.Sprintf(callBackImportUrl, importReq.db, importReq.schema, importReq.table), result)
		w.WriteHeader(http.StatusOK)
	}
}
