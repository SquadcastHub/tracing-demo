package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

var (
	version              = "unknown"
	envKeyBackServiceURL = "BACKEND_SVC_URL"
)

func main() {
	addr := os.Getenv("FRONTEND_APP_ADDR")
	r := chi.NewRouter()
	r.Get("/counter/{keyspace}", func(w http.ResponseWriter, r *http.Request) {
		keyspace := chi.URLParam(r, "keyspace")
		count, err := getCount(keyspace)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write([]byte(fmt.Sprintf("Counter for %s: %s", keyspace, count)))
	})
	// execute server
	log.Printf("front service is listening on %v", addr)
	http.ListenAndServe(addr, r)
}

func getCount(keyspace string) (string, error) {
	requestURL := fmt.Sprintf("http://%s/%s", os.Getenv(envKeyBackServiceURL), keyspace)
	resp, err := http.Get(requestURL)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)

	return sb, nil
}
