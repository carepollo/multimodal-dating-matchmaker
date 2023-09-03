package handlers

import (
	"log"
	"net/http"

	"github.com/carepollo/multimodal-dating-matchmaker/auth/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	list := data.GetProducts()
	err := list.ToJSON(w)
	if err != nil {
		http.Error(w, "failed to marshal json", http.StatusInternalServerError)
	}
}


func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	
}