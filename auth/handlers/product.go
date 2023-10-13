package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

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

	if r.Method == http.MethodPut {
		c := regexp.MustCompile(`/([0-9]+)`)
		g := c.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 1 {
			http.Error(w, "invalid url, not found value", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			http.Error(w, "invalid url, not compiled value", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, "unable to convert to number", http.StatusBadRequest)
			return
		}

		p.l.Println("got id", id)
		p.updateProduct(id, w, r)

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
	p.l.Println("added product in post")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to umarshal JSON", http.StatusBadRequest)
	}

	p.l.Printf("product: %v", prod)
	data.AddProduct(prod)
}

func (p *Products) updateProduct(id int, w http.ResponseWriter, r *http.Request) {
	p.l.Println("modified product in put")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to umarshal JSON", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "product not valid", http.StatusNotFound)
		return
	}

}
