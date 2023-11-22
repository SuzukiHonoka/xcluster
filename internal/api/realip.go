package api

import (
	"net"
	"net/http"
	"strings"
)

// inspiration from https://gist.github.com/miguelmota/7b765edff00dc676215d6174f3f30216

func RealIP(r *http.Request) string {
	ips := r.Header.Get("X-Forwarded-For")
	splitIps := strings.Split(ips, ",")
	if len(splitIps) > 0 {
		// get last IP in list since ELB prepends other user defined IPs, meaning the last one is the actual client IP.
		netIP := net.ParseIP(splitIps[len(splitIps)-1])
		if netIP != nil {
			return netIP.String()
		}
	}
	return strings.Split(r.RemoteAddr, ":")[0]
}
