package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/s0a1qq/booking-mess/internal/config"
)

var app *config.AppConfig

// NewHelpers sets up config for helpers
func NewHelpers(a *config.AppConfig) {
	app = a
}

func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("ClientError with status off", status)
	http.Error(w, http.StatusText(status), status)
}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func IsAuthenticated(r *http.Request) bool {
	exists := app.Session.Exists(r.Context(), "user_id")
	return exists
}
