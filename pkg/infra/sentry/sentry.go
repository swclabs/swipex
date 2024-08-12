// Package sentry connect to sentry server
package sentry

import (
	"fmt"
	"swclabs/swix/internal/config"

	"github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
)

// Init sentry
func Init() {
	if config.StageStatus != "dev" {
		if err := sentry.Init(sentry.ClientOptions{
			Dsn:           config.SentryDsn,
			EnableTracing: true,
			// Set TracesSampleRate to 1.0 to capture 100%
			// of transactions for performance monitoring.
			// We recommend adjusting this value in production,
			TracesSampleRate: 1.0,
			AttachStacktrace: true,
		}); err != nil {
			fmt.Printf("Sentry initialization failed: %v", err)
		}
	}
}

// CaptureMessage capture message
func CaptureMessage(c echo.Context, message string) {
	if config.StageStatus != "dev" {
		if hub := sentryecho.GetHubFromContext(c); hub != nil {
			hub.WithScope(func(_ *sentry.Scope) {
				// scope.SetExtra("unwantedQuery", "someQueryDataMaybe")
				hub.CaptureMessage(message)
			})
		}
	}
}
