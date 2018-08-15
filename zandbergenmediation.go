package main

import (
	_ "fmt"
	"net/http"
)

func NewServeViewFunc(view string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(view))
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/html/index.html", 301)
	// fmt.Fprint(w, "Hello, world!")
}

func RegisterHandlers() {
	indexHandlerFunc := NewServeViewFunc(indexView)
	http.HandleFunc("/index.html", indexHandlerFunc)
	http.HandleFunc("/contact.html", NewServeViewFunc(contactView))
	http.HandleFunc("/mediation.html", NewServeViewFunc(mediationView))
	http.HandleFunc("/nmi.html", NewServeViewFunc(nmiView))
	http.HandleFunc("/uwmediator.html", NewServeViewFunc(uwMediatorView))
	http.HandleFunc("/sitemap.xml", handleSitemap)
	http.HandleFunc("/", indexHandlerFunc)
}
