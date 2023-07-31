package handlers

import (
	"log"
	"net/http"
)

type Login struct {
	l *log.Logger
}

func NewLogin(l *log.Logger) *Login {
	return &Login{l}
}

func (login *Login) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bye lol"))
}
