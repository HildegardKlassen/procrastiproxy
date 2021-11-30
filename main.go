package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8000", nil)

}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Sorry Hila, aber diese Seite ist für dich gesperrt!!")
	w.WriteHeader(403)
}
