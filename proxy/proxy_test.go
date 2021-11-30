package proxy

import (
	"net/http"
	"net/url"
	"testing"
)

func Test_Proxy(t *testing.T) {

	go func() { Run() }()

	pURL, _ := url.Parse("http://localhost:8000")

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(pURL),
		},
	}

	res, err := client.Get("http://www.google.com")

	if err != nil {
		t.Error(err)

	}

	if res.StatusCode != 403 {
		t.Error(err)
	}

}
