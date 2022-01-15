package reverse_proxy

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ServeReverseProxy(target string, w http.ResponseWriter, r *http.Request) {
	url, _ := url.Parse(target)
	fmt.Println(url)
	proxy := httputil.NewSingleHostReverseProxy(url)
	fmt.Println(proxy)
	proxy.ServeHTTP(w, r)
}
