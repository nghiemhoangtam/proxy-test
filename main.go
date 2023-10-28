package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func reverseProxy(targetURL *url.URL) http.Handler {
	return &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL = targetURL
			req.Host = targetURL.Host
		},
	}
}

func main() {
	targetURL, _ := url.Parse("https://medium.com")
	proxy := reverseProxy(targetURL)
	http.Handle("/", proxy)
	http.ListenAndServe(":8080", nil)
}
