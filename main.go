// The end-to-end suite's fixture repository (steloit/cloud #77): a Web
// Service the platform builds from this repository on its build Cluster and
// deploys. It listens on PORT and answers /healthz, the suite's health path,
// and / with a line the suite can expect.
package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte("ok\n"))
	})
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte("steloit e2e fixture\n"))
	})
	srv := &http.Server{Addr: ":" + port, Handler: mux, ReadHeaderTimeout: 10 * time.Second}
	log.Printf("listening on :%s", port)
	log.Fatal(srv.ListenAndServe())
}
