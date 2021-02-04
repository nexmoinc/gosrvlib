package httpserver

import (
	"net/http"

	"github.com/nexmoinc/gosrvlib/pkg/logging"
	"github.com/nexmoinc/gosrvlib/pkg/traceid"
	"github.com/nexmoinc/gosrvlib/pkg/uidc"
	"go.uber.org/zap"
)

// RequestInjectHandler wraps all incoming requests and injects a logger in the request scoped context.
func RequestInjectHandler(rootLogger *zap.Logger, traceIDHeaderName string, next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		reqID := traceid.FromHTTPRequestHeader(r, traceIDHeaderName, uidc.NewID128())

		reqLog := rootLogger.With(
			zap.String("traceid", reqID),
			zap.String("request_method", r.Method),
			zap.String("request_path", r.URL.Path),
			zap.String("request_query", r.URL.RawQuery),
			zap.String("request_uri", r.RequestURI),
			zap.String("request_useragent", r.UserAgent()),
			zap.String("remote_ip", r.RemoteAddr),
		)

		ctx = logging.WithLogger(ctx, reqLog)
		ctx = traceid.NewContext(ctx, reqID)

		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
