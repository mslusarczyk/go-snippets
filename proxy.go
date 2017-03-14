package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/gorilla/handlers"
)

const TargetUrl = "http://localhost:8000"
const ListeningPort = "8888"

func main() {
	// intercept call
	http.Handle("/web-hook", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(ProxyCall)))

	// catch-all
	http.Handle("/", http.NotFoundHandler())
	log.Println("Starting up..")
	http.ListenAndServe(":"+ListeningPort, nil)
	log.Println("Done.")
}

func ProxyCall(w http.ResponseWriter, r *http.Request) {
	log.Println("Proxy call for web-hook")
	u, err := url.Parse(TargetUrl)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ServeHTTP(w, r)
}
