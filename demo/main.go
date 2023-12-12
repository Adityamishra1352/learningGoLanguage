package main

import (
	"fmt"
	"net/http"
)

func helloWorldPage(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "Hello world")
	case "/a":
		fmt.Fprint(w, "Aditya")
	case "/b":
		fmt.Fprint(w, "Ninnu")
	case "/c":
		fmt.Fprint(w, "Kittu")
	default:
		fmt.Fprint(w, "Bit fat error")
	}

}
func main() {
	http.HandleFunc("/", helloWorldPage)
	http.ListenAndServe("", nil)
}
