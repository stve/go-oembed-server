package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/widnyana/oembed"
)

// Error - default error message type
type Error struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

const defaultPort string = "8000"

func extract(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	encoder := json.NewEncoder(w)

	providers, err := ioutil.ReadFile("providers.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeError(encoder, "Unable to load supported providers")
		return
	}

	service := oembed.NewOembed()
	service.ParseProviders(bytes.NewReader(providers))

	url := r.URL.Query().Get("url")
	if isEmpty(url) {
		w.WriteHeader(http.StatusBadRequest)
		writeError(encoder, "Invalid URL")
		return
	}

	item := service.FindItem(url)
	if item == nil {
		w.WriteHeader(http.StatusBadRequest)
		writeError(encoder, "Unsupported oembed provider")
		return
	}

	info, err := item.FetchOembed(url, nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeError(encoder, "Unsupported oembed provider")
		return
	}

	resp, err := info.ToJSON()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeError(encoder, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)

	return
}

func writeError(encoder *json.Encoder, message string) {
	if err := encoder.Encode(Error{Type: "error", Message: message}); err != nil {
		panic(err)
	}
}

func isEmpty(str string) bool {
	return (len(str) == 0)
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = defaultPort
	}

	http.HandleFunc("/", extract)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
