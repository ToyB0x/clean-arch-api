package middleware

import (
	"net/http"

	"github.com/toaru/clean-arch-api/config"
)

func LimitSizeByContentLengthHeader(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.ContentLength > config.Configs.MaxRequestSize {
			_ = BuildErrorResponse(w, "content length over")
			return // http.HandlerFunc(fn) を返さずに処理を終了する
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
