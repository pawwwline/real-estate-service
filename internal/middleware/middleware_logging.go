package middleware

import(
	"log/slog"
	"net/http"
)

func LoggerMiddleware(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Info("Request received",
				"method", r.Method,
				"url", r.URL.String(),
				"remote_addr", r.RemoteAddr,
			)

			next.ServeHTTP(w, r)
			
			logger.Info("Response sent",
				"method", r.Method,
				"url", r.URL.String(),
				"remote_addr", r.RemoteAddr,
			)
		})
	}
}