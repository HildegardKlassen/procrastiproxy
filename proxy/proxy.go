package proxy

import (
	"net/http"
)

func Run() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusForbidden)
}
