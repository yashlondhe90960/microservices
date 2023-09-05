package handlers

import(
	"fmt"
	"log"
	"net/http"


)

type Goodbye struct{
	l *log.Logger
}

func NewGoodBye(l *log.Logger) *Goodbye{
	return &Goodbye{l}
}


func (h * Goodbye) ServerHTTP(rw http.ResponseWriter, r *http.Request){
	h.l.Println("handle Goodbye request")

	//writing he response
	fmt.Fprintf(rw, "Goodbye")
}