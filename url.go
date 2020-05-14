package main

import (
	"fmt"
	"strings"
)

const (
	HTTP  = "http://"
	HTTPS = "https://"
)

func hasProtocolPrefix(url string) bool {
	return strings.HasPrefix(url, HTTP) || strings.HasPrefix(url, HTTPS)
}

func compose(url string) string {
	if hasProtocolPrefix(url) {
		// return unmodified url
		return url
	} else {
		// fallback to http
		return fmt.Sprintf("%s%s", HTTP, url)
	}
}
