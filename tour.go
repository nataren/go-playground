package main

import (
	"net/http"
	"fmt"
)

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", s)
}

func (st *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v%v %v", st.Greeting, st.Punct, st.Who)
}

func main() {
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct { "Hello", ":", "Gophers!" })
	http.ListenAndServe("localhost:4000", nil)
}