package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Ping struct {
	l *log.Logger
}

func NewPing(l *log.Logger) *Ping {
	return &Ping{l}
}

func (ping *Ping) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Not valid response", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Hello %s", name)
}
