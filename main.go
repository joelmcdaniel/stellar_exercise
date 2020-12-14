package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var host string
var port string
var key string
var mSResp map[string]snippetResp

type snippetReq struct {
	Name      string
	ExpiresIn int
	Snippet   string
}

type snippetResp struct {
	URL       string `json:"url"`
	Name      string `json:"name"`
	ExpiresAt string `json:"expires_at"`
	Snippet   string `json:"snippet"`
}

func init() {
	gotenv.Load()
	host = os.Getenv("HOST")
	port = os.Getenv("PORT")

	key = "recipe"
	mSResp = make(map[string]snippetResp)
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/snippets", snippetsHandler)
	r.HandleFunc(fmt.Sprintf("/snippets/{%s}", key), snippetsHandler)

	log.Fatal(http.ListenAndServe(":"+port, r))
}

// SnippetsHandler ...
func snippetsHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	if key, ok := vars[key]; ok {
		// /snippets/{recipe} route hit

		if s, ok := mSResp[key]; ok {
			// recipe snippet already exists

			t1 := time.Now()
			t2, err := time.Parse(time.RFC3339, s.ExpiresAt)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if t1.After(t2) {
				// recipe snippet expired
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
			// recipe snippet doesn't exist
			w.WriteHeader(http.StatusNotFound)
		}

	} else {
		// /snippets route hit

		var sReq snippetReq
		err := json.NewDecoder(r.Body).Decode(&sReq)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		t1 := time.Now()
		t2 := t1.Add(time.Second * 30)

		sResp := snippetResp{
			URL:       fmt.Sprintf("https://%s:%s/snippets/%s", host, port, sReq.Name),
			Name:      sReq.Name,
			ExpiresAt: t2.Format(time.RFC3339),
			Snippet:   sReq.Snippet,
		}

		mSResp[sReq.Name] = sResp

		resp, err := json.Marshal(sResp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(resp))
	}

}
