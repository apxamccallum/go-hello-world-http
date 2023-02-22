package main

import (
	"fmt"
	"net/http"
	"os"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	body := os.Getenv("WEBBODY")
	fmt.Fprint(w, body)
}

func main() {
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe("0.0.0.0:80", nil)
}
