package main

import (
	"net/http"
	"os"
	"path/filepath"
)

type static struct {
	dir string
	h   http.Handler
}

func (s static) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fname := filepath.Join(s.dir, r.URL.Path)

	_, err := os.Stat(fname)
	if err != nil {
		s.h.ServeHTTP(w, r)
		return
	}

	server := http.FileServer(http.Dir(s.dir))
	server.ServeHTTP(w, r)
}

func Static(dir string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return static{dir, h}
	}
}
