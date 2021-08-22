package middleware

import (
	"context"
	"net/http"
)

func GetCloudFlareIp(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		CF_IP := r.Header.Get("CF-Connecting-IP")
		ctx := context.WithValue(r.Context(), "cf_ip", CF_IP)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
