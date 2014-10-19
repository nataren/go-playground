package main 

import (
	"fmt"
	"net/http"
)

type Hello struct { }

func (h Hello) ServeHTTP(w http.ResponseWriter,	r *http.Request) {
	fmt.Fprint(w, "The request: ", r)
}

func main() {
	var h Hello
	http.ListenAndServe("localhost:4040", h)
}