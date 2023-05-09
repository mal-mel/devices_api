package api

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func LoggerMiddleware(log *logrus.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			newLogger := log.WithFields(map[string]interface{}{
				"method": req.Method,
				"url":    req.URL.String(),
				"host":   req.Host,
			})

			newLogger.Info("request")

			next.ServeHTTP(res, req)
		})
	}
}
