package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Payload struct {
	When     int64
	response string
}

func DoCallbackRequest(url string, msg string) {
	p := Payload{response: msg, When: time.Now().Unix()}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(p)
	_, err := http.Post(url, "application/json;charset=utf-8", b)
	if err != nil {
		log.Println(fmt.Sprintf("ERROR CALLING CALLBACK URL [%v] with message [%v]: %v", url, msg, err))
	}
}
