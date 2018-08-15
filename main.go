// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
)

func main() {
	// http.HandleFunc("/", handle)
	RegisterHandlers()
	appengine.Main()
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, brave new world!")
}

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
