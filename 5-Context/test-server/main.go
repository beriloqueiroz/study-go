package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	select {
	case <-time.After(10 * time.Second):
		log.Println("sucesso ao processar requisição")
		w.Write([]byte("ok"))
	case <-ctx.Done():
		log.Println("requisição cancelada pelo usuário")
	}
}
