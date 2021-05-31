package middleware

import "net/http"

type HandleFunc func(w http.ResponseWriter, r *http.Request)
type MiddlewareFunc func(mw *middleware)
type Middleware = middleware
