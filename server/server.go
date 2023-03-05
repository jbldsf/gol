package server

import (
	"fmt"
	"gol/config"
	"gol/service"
	"net/http"
	"strings"
)

type handler struct{}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router(r, func(statusCode int, response []byte) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		w.Write(response)
	})
}

func router(r *http.Request, callback func(statusCode int, response []byte)) {
	switch strings.Trim(r.URL.Path, "/") {
	case "countries":
		service.Country(r, callback)
	default:
		callback(http.StatusNotFound, []byte("{}"))
	}
}

func HTTP() {
	fmt.Printf("HTTP server running on port %d", config.HTTPPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.HTTPPort), handler{})
	if err != nil {
		println(err.Error())
	}
}

func HTTPS() {
	fmt.Printf("HTTPS server running on port %d", config.HTTPSPort)
	err := http.ListenAndServeTLS(fmt.Sprintf(":%d", config.HTTPSPort), "", "", handler{})
	if err != nil {
		println(err.Error())
	}
}
