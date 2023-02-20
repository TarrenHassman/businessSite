package main

import (
	"net/http"
)

func main() {
	static := http.FileServer(http.Dir("./web"))
	http.Handle("/", static)
	http.ListenAndServe(":8080", nil)
}
