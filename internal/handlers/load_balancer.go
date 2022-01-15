package handlers

import (
	"net/http"

	"github.com/SantiagoBedoya/go-proxy-server/internal/logger"
	"github.com/SantiagoBedoya/go-proxy-server/internal/reverse_proxy"
)

var serverCount = 0

const (
	SERVER1 = "https://jsonplaceholder.typicode.com/todos"
	SERVER2 = "https://www.facebook.com"
	SERVER3 = "https://www.yahoo.com"
)

func LoadBalancer(w http.ResponseWriter, r *http.Request) {
	url := getProxyUrl()
	logger.LogRequestPayload(url)
	reverse_proxy.ServeReverseProxy(url, w, r)
}

func getProxyUrl() string {
	var servers = []string{SERVER1, SERVER2, SERVER3}
	server := servers[serverCount]
	serverCount++

	if serverCount >= len(servers) {
		serverCount = 0
	}

	return server
}
