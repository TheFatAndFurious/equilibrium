package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func newProxy(targetHost string) (*httputil.ReverseProxy, error) {
	forwardUrl, err := url.Parse(targetHost)
	if err != nil {
		return nil, err
	}
	return httputil.NewSingleHostReverseProxy(forwardUrl), nil
}

func ProxyRequestHandler(proxy *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		proxy.ServeHTTP(res, req)
	}
}
func main() {
	proxy, err := newProxy("http://localhost:8090")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", ProxyRequestHandler(proxy))
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
