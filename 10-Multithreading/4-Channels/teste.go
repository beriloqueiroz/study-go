package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	ch := make(chan string)
	go sendEvent(ch)
	listener(ch)

	http.HandleFunc("/", handler)

	http.ListenAndServe(":8081", nil)
}

func middlewareEvent(ch chan string) http.Handler {

}

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
}

func sendEvent(ch chan string, method string) {
	for i := 0; i < 10; i++ {
		ch <- "task " + strconv.Itoa(i) + "method: " + method
	}
	close(ch)
}

func listener(ch chan string) {
	for event := range ch {
		fmt.Printf("Evento: %s, disparado\n", event)
	}
}
