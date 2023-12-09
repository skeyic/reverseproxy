package router

import (
	"net/http"
	"net/http/httputil"
	url2 "net/url"
)

var (
	TheRouter = &Router{}
)

type Router struct {
}

func (r *Router) Do(w http.ResponseWriter, req *http.Request) {
	if req.RequestURI == "/baidu" {
		url, _ := url2.Parse("http://www.baidu.com")
		proxy := httputil.NewSingleHostReverseProxy(url)
		req.Host = "http://www.baidu.com"
		req.URL.Path = "/"
		proxy.ServeHTTP(w, req)
		return
	}
	w.Write([]byte("Not implemented"))
}
