package handlers

import (
	"fmt"
	"log"
	"net/http"
	"io"

)

type Hello struct {
	l *log.Logger

}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// ServerHTTP implements the http.Server interface
// https://golang.org/pkg/net/http/#Handler


func (h *Hello) SeverHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle Hello requests")

	//read a body
	b, err := io.ReadAll(r.Body)
	if err != nil {
		h.l.Println("Error reading response body", err)

		http.Error(rw, "Unable to read request body", http.StatusBadRequest)
		return
	}	

	//write the response body
	fmt.Fprintf(rw, "Hello %s", b)




}