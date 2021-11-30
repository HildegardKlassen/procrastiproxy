package proxy

import (
	"net/http"
	"net/url"
	"sync"
	"testing"
)

var testProxy *Proxy
var once sync.Once

func TestMain(m *testing.M) {
	once.Do(
		func() {
			testProxy = NewProxy("localhost", "8000")
			go func() {
				testProxy.Run()
			}()
		})
}

func Test_Proxy(t *testing.T) {

	req, err := http.NewRequest("GET", "https://www.google.de", nil)
	if err != nil {
		t.Error(err)
	}

	testUrl := url.URL{
		Host:   req.URL.Host,
		Path:   req.URL.Path,
		Scheme: req.URL.Scheme,
	}

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(&testUrl),
		},
	}

	res, err := client.Get("https://www.google.de")

	if err != nil {
		t.Error(err)

	}

	if res.StatusCode != 403 {
		t.Error(err)
	}

	//body, err = ioutil.ReadAll(res.Body)

}
