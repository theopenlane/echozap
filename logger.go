package echozap

import (
	"fmt"
	"time"

	echo "github.com/theopenlane/echox"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ZapLogger is a middleware and zap to provide an "access log" like logging for each request.
func ZapLogger(log *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)
			if err != nil {
				c.Error(err)
			}

			req := c.Request()
			res := c.Response()

			fields := []zapcore.Field{
				zap.String("remote_ip", c.RealIP()),
				zap.String("latency", time.Since(start).String()),
				zap.String("host", req.Host),
				zap.String("request", fmt.Sprintf("%s %s", req.Method, req.RequestURI)),
				zap.Int("status", res.Status),
				zap.Int64("size", res.Size),
				zap.String("user_agent", req.UserAgent()),
			}

			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
			}

			fields = append(fields, zap.String("request_id", id))

			n := res.Status

			switch {
			case n >= 500: // nolint: mnd
				log.With(zap.Error(err)).Error("Server error", fields...)
			case n >= 400: // nolint: mnd
				log.With(zap.Error(err)).Warn("Client error", fields...)
			case n >= 300: // nolint: mnd
				log.Info("Redirection", fields...)
			default:
				log.Info("Success", fields...)
			}

			return nil
		}
	}
}
