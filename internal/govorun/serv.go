package govorun

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

var (
	sayPath         = "/say"
	listenPath      = "/listen"
	jsonApplication = "application/json"
)

type SayBody struct {
	Word string `json:"word"`
}

func (gov *Govorun) startServer() {
	log.Fatal("HTTP server error: ", http.ListenAndServe("localhost:5000", gov))
}
func (gov *Govorun) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	if isSay(req) {
		var body SayBody
		err := json.NewDecoder(req.Body).Decode(&body)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		gov.UpdateMessage(body.Word)
		rw.WriteHeader(201)
		return
	}

	if isListen(req) {
		flusher, ok := rw.(http.Flusher)

		if !ok {
			http.Error(rw, "Streaming unsupported!", http.StatusInternalServerError)
			return
		}

		rw.Header().Set("Content-Type", "text/event-stream")
		rw.Header().Set("Cache-Control", "no-cache")
		rw.Header().Set("Connection", "keep-alive")
		rw.Header().Set("Access-Control-Allow-Origin", "*")

		timestamp := time.Now().Unix()
		name := strconv.FormatInt(timestamp, 10)
		ch := make(chan string)

		un := gov.Subscribe(name, ch)

		done := req.Context().Done()

		var msg string

		for {
			select {
			case msg = <-ch:
				_, err := fmt.Fprintf(rw, "%s channel %s\n", name, msg)
				if err != nil {
					return
				}
				flusher.Flush()
			case <-done:
				un()
				close(ch)
				return
			}
		}
	}

	log.Println(req.Method, req.URL)
}

func isSay(req *http.Request) bool {
	return req.Method == http.MethodPost &&
		req.URL.Path == sayPath &&
		req.Header.Get("Content-Type") == jsonApplication
}

func isListen(req *http.Request) bool {
	return req.Method == http.MethodGet && req.URL.Path == listenPath
}
