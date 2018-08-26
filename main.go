// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func main() {
	// http.HandleFunc("/", handle)
	RegisterHandlers()
	appengine.Main()
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, brave new world!")
}

type prefixHandler struct {
	from string
	to   string
	h    http.Handler
}

func (ph *prefixHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Debugf(ctx, "prefix. Host is: %s", r.Host)
	// Split port and host.
	s := strings.Split(r.Host, ":")
	log.Debugf(ctx, "prefix. elements: %d", len(s))
	log.Debugf(ctx, "prefix. s: %v", s)

	// Test if hosts does not start with from.
	if strings.ToLower(s[0]) != strings.ToLower(ph.from) {
		// Call original handler.
		log.Debugf(ctx, "prefix. calling original handler")
		ph.h.ServeHTTP(w, r)
		return
	}

	// Build the url to redirect
	s[0] = ph.to
	log.Debugf(ctx, "prefix. modified s: %v", s)
	u := *r.URL
	u.Host = strings.Join(s, ":")
	log.Debugf(ctx, "prefix. modified url: %s", u.String())
	http.Redirect(w, r, u.String(), http.StatusMovedPermanently)
}

func HostPrefixHandlerFunc(prefix string, handlerFunc http.HandlerFunc) http.HandlerFunc {
	h := handlerFunc
	p := prefix
	return func(w http.ResponseWriter, r *http.Request) {
		// Test if hosts starts with prefix.
		host := r.URL.Host
		if strings.HasPrefix(host, p) {
			// Call original handler.
			h.ServeHTTP(w, r)
		} else {
			// Build the url to redirect
			u := *r.URL
			u.Host = p + u.Host
			http.Redirect(w, r, u.String(), http.StatusMovedPermanently)
		}
	}
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
	// Create a servemutex for the routing.
	sm := http.NewServeMux()
	indexHandlerFunc := NewServeViewFunc(indexView)
	sm.HandleFunc("/index.html", indexHandlerFunc)
	sm.HandleFunc("/contact.html", NewServeViewFunc(contactView))
	sm.HandleFunc("/mediation.html", NewServeViewFunc(mediationView))
	sm.HandleFunc("/nmi.html", NewServeViewFunc(nmiView))
	sm.HandleFunc("/uwmediator.html", NewServeViewFunc(uwMediatorView))
	sm.HandleFunc("/sitemap.xml", handleSitemap)
	sm.HandleFunc("/", indexHandlerFunc)
	// Prefix handler handles the requests before sm.
	ph := &prefixHandler{
		from: os.Getenv("URL_FROM"),
		to:   os.Getenv("URL_TO"),
		h:    sm,
	}
	// Regiser the handler for all paths.
	http.Handle("/", ph)
}
