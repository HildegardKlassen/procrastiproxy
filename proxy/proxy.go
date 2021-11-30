package proxy

import (
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/gorilla/mux"
)

func RunAllForbidden() {
	http.ListenAndServe("localhost:9996", forbiddenRouter())
}

func RunAllAllowed() {
	http.ListenAndServe("localhost:9997", allowedRouter())
}

func RunBlockList() {
	http.ListenAndServe("localhost:9998", blockListRouter())
}

func forbiddenHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusForbidden)
}

func allowedHandler(w http.ResponseWriter, r *http.Request) {
	rp := httputil.NewSingleHostReverseProxy(r.URL)
	rp.ServeHTTP(w, r)
}

func blockListHandler(w http.ResponseWriter, r *http.Request) {
	blockList := parseBlockList()
	for _, b := range blockList {
		if strings.Contains(b, r.URL.Host) {
			if strings.Contains(b, r.URL.Path) || (r.URL.Path == "/") {
				w.WriteHeader(http.StatusForbidden)
				return
			}
		}
	}
	rp := httputil.NewSingleHostReverseProxy(r.URL)
	rp.Director = func(req *http.Request) {
		req.URL = r.URL
	}
	//rp.Transport = &tls.Config{InsecureSkipVerify: true}
	rp.ServeHTTP(w, r)
}

func forbiddenRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", forbiddenHandler)
	return r
}

func allowedRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", allowedHandler)
	return r
}

func blockListRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", blockListHandler)
	r.PathPrefix("/").HandlerFunc(blockListHandler)
	return r
}

func parseBlockList() []string {
	content, err := ioutil.ReadFile("blocklist")
	if err != nil {
		return nil
	}
	lines := strings.Split(string(content), ",")
	return lines
}
