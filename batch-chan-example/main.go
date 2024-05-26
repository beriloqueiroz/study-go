package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

var batch []string
var ch chan string

func main() {
	ch = make(chan string, 5)
	batch = []string{}

	go func() {
		defer close(ch)
		for {
			select {
			case input, ok := <-ch:
				if !ok {
					if len(batch) > 0 {
						fmt.Printf("Houve erro, enviando residual %s", batch)
					}
					return
				}
				batch = append(batch, input)
				if len(batch) >= 5 {
					fmt.Println("insert all ", batch)
					batch = []string{}
				}
			}
		}
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /lance", handler)
	http.ListenAndServe(":3333", mux)
}

func handler(w http.ResponseWriter, r *http.Request) {
	input := uuid.New().String()
	ch <- input
	w.WriteHeader(200)
}
