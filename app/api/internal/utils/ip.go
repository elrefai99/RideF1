package utils

import "net/http"

func GetIP(r *http.Request) string {

	if ip := r.Header.Get("cf-connecting-ip"); ip != "" {
		return ip
	}

	if ip := r.Header.Get("x-real-ip"); ip != "" {
		return ip
	}

	if ip := r.Header.Get("x-forwarded-for"); ip != "" {
		return ip
	}

	return r.RemoteAddr
}
