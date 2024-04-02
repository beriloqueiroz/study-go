package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("https://api.github.com/users/beriloqueiroz")
	if err != nil {
		panic("erro 1")
	}
	defer req.Body.Close()
	result, err := io.ReadAll(req.Body)
	if err != nil {
		panic("erro 2")
	}
	fmt.Print(string(result))
}
