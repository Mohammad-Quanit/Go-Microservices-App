package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHttp(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello from Hello logger")
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Oops", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(rw, "Hello %s from response writer", b)
}