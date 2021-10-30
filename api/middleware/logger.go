package middleware

// Logger middleware logs the information about each HTTP request.

import (
	"fmt"
	"net/http"
	"time"

	echo "github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// Penformatan pencatatan sistem
func MakeLogEntry(c echo.Context) *log.Entry {
	start := time.Now()

	if c == nil {
		return log.WithFields(log.Fields{
			"latency_ns": time.Since(start).Nanoseconds(),
			"group":      "system",
		})
	}

	return log.WithFields(log.Fields{
		"group":      "request",
		"method":     c.Request().Method,
		"uri":        c.Request().RequestURI,
		"remoteaddr": c.Request().RemoteAddr,
		"latency_ns": time.Since(start).Nanoseconds(),
	})
}

func MiddlewareLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		MakeLogEntry(c).Info("Incoming Connection")
		return next(c)
	}
}

func ErrorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if ok {
		report.Message = fmt.Sprintf("http error %d - %v", report.Code, report.Message)
	} else {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	MakeLogEntry(c).Error(report.Message)
	c.HTML(report.Code, report.Message.(string))
}
