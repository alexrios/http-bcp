package main

import (
	"bytes"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

type Payload struct {
	When     int64
	response string
}

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func DoCallbackRequest(url string, msg string) {
	p := Payload{response: msg, When: time.Now().Unix()}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(p)
	callbackLogger := log.WithFields(log.Fields{
		"phase":   "runtime",
		"event":   "calling callback service",
		"url":     url,
		"payload": p,
	})
	_, err := http.Post(url, "application/json;charset=utf-8", b)
	if err != nil {
		callbackLogger.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Call failed")
	} else {
		callbackLogger.Info("Succesful call")
	}
}
