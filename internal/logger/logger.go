package logger

import "log"

func LogRequestPayload(proxyUrl string) {
	log.Printf("proxy_url: %s\n", proxyUrl)
}
