package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var host string
var mS map[string]snippet

type snippet struct {
	URL     string `json:"url"`
	Name    string `json:"name"`
	Expires string `json:"expires_at"`
	Snippet string `json:"snippet"`
}

func init() {

	host = "example.com"
	mS = make(map[string]snippet)
}

func main() {

	crt := flag.String("cert", "", "the path the the cert file")
	key := flag.String("key", "", "the path to the key file")

	flag.Parse()

	r := mux.NewRouter()
	s := r.Host(host).PathPrefix("/snippets").Subrouter()
	s.HandleFunc("", postHandler).Methods("POST")
	s.HandleFunc("/{recipe}", getHandler).Methods("GET")

	err := http.ListenAndServeTLS(":443", *crt, *key, r)
	if err != nil {
		log.Fatal("ListenAndServeTLS: ", err)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {

	var sReq snippet
	err := json.NewDecoder(r.Body).Decode(&sReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	t1 := time.Now()
	t2 := t1.Add(time.Second * 30)

	sResp := snippet{
		URL:     fmt.Sprintf("https://%s/snippets/%s", host, sReq.Name),
		Name:    sReq.Name,
		Expires: t2.Format(time.RFC3339),
		Snippet: sReq.Snippet,
	}

	mS[sReq.Name] = sResp

	resp, err := json.Marshal(sResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(resp))
}

func getHandler(w http.ResponseWriter, r *http.Request) {

	for k := range mux.Vars(r) {
		if s, ok := mS[k]; ok {
			// recipe snippet already exists in map
			t1 := time.Now()
			t2, err := time.Parse(time.RFC3339, s.Expires)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if t1.After(t2) {
				// recipe snippet expired
				delete(mS, s.Name)
				w.WriteHeader(http.StatusNotFound)
			} else {
				resp, err := json.Marshal(s)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(resp))
			}
		} else {
			// recipe snippet doesn't exist in map
			w.WriteHeader(http.StatusNotFound)
		}
	}
}
