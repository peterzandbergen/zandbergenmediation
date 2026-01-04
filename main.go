// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"embed"
	"io"
	"io/fs"
	"log/slog"
	"mime"
	"net/http"
	"os"
	"time"
)

func initLogger(w io.Writer, application string) {
	lh := slog.NewTextHandler(w, &slog.HandlerOptions{Level: slog.LevelInfo})
	logger := slog.New(lh)
	slog.SetDefault(logger.With("application", application))
}

func run(ctx context.Context, getenv func(string) string) {
	log := slog.Default()
	registerHandlers(ctx)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Info("using default port", "port", port)
	}

	log.Info("starting ListenAndServe", "port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Error("ListenAndServe failed", "err", err)
	}
}

func main() {
	ctx := context.Background()
	initLogger(os.Stdout, "zandbergenmediation")
	run(ctx, os.Getenv)
}

//go:embed html/*
var htmlRoot embed.FS

func init() {
	// set mime type for .css files
	mime.AddExtensionType(".css", "text/css; charset=utf-8")
}

func slogHandler(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		defer func(start time.Time) {
			logger.InfoContext(ctx, "served request", "method", r.Method, "url", r.URL.String(), "remote_addr", r.RemoteAddr, "duration_ms", time.Since(start).Milliseconds())
		}(time.Now())

		next.ServeHTTP(w, r)
	})
}

func registerHandlers(ctx context.Context) {
	root, err := fs.Sub(htmlRoot, "html")
	if err != nil {
		slog.ErrorContext(ctx, "cannot open sub html", "err", err)
	}

	http.Handle("/", slogHandler(slog.Default(), http.FileServer(http.FS(root))))
}
