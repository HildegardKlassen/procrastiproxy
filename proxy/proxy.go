package proxy

import (
	"fmt"
	"net/http"
)

func NewProxy(url, port string) *Proxy {
	return &Proxy{
		url:  url,
		port: port}

}

func (p *Proxy) Run() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf("%s:%s", p.url, p.port), nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
	w.WriteHeader(200)
}

type Proxy struct {
	url  string
	port string
}
