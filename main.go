package main

import (
	"fmt"
	"net/http"
	"os"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

type nacho struct{}

func (na nacho) Write(p []byte) (n int, err error) {
	fmt.Println(fmt.Sprintf("%s", p))
	return len(p), nil
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Server is running on port 3000")
	fmt.Fprintln(os.Stdout, "Hello world")
	fmt.Fprintln(nacho{}, "Hello world")
	http.ListenAndServe(":3000", nil)
}
