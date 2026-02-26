package middleware

import "net/http"

func Admin(next http.Handler) http.Handler { return next }
