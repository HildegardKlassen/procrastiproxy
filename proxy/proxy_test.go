package proxy

import (
	"net/http"
	"net/url"
	"testing"
)

const (
	forbiddenProxyURL = "http://localhost:9996"
	allowedProxyURL   = "http://localhost:9997"
	blockListProxyURL = "http://localhost:9998"
)

func Test_ProxyEverythingAllowed(t *testing.T) {

	go RunAllAllowed()

	pURL, _ := url.Parse(allowedProxyURL)

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(pURL),
		},
	}

	res, err := client.Get("http://www.google.com")

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Error(err)
	}

}

func Test_ProxyEverythingForbidden(t *testing.T) {

	go RunAllForbidden()

	pURL, _ := url.Parse(forbiddenProxyURL)

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

func Test_ProxyBlockList(t *testing.T) {

	go RunBlockList()

	pURL, _ := url.Parse(blockListProxyURL)

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
