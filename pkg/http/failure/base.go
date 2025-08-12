package failure

import "net/http"

type App struct {
	Code    int
	Message string
	Err     error
}

func NewApp(code int, message string, err error) *App {
	return &App{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func NewInternal(err error) *App {
	return &App{
		Code:    http.StatusInternalServerError,
		Message: "Terjadi kesalahan tak terduga pada server.",
		Err:     err,
	}
}
