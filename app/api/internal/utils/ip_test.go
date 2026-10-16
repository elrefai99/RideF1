package utils

import (
	"net/http/httptest"
	"testing"
)

func TestIp(t *testing.T) {
	tests := []struct {
		name       string
		header     string
		value      string
		remoteAddr string
		want       string
	}{
		{"Cloudflare header", "cf-connecting-ip", "1.2.3.4", "", "1.2.3.4"},
		{"Real IP header", "x-real-ip", "5.6.7.8", "", "5.6.7.8"},
		{"Forwarded header", "x-forwarded-for", "9.10.11.12", "", "9.10.11.12"},
		{"Remote address fallback", "", "", "127.0.0.1:1234", "127.0.0.1:1234"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			if tt.header != "" {
				req.Header.Set(tt.header, tt.value)
			}
			req.RemoteAddr = tt.remoteAddr

			if got := GetIP(req); got != tt.want {
				t.Errorf("GetIP() = %q, want %q", got, tt.want)
			}
		})
	}
}
