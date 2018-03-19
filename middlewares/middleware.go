package middlewares

import "github.com/gorilla/mux"

func RegisterMiddlewares(r *mux.Router) {
	var (
		authentication = Authentication{WhitelistedRoutes: []string{"/login"}}
	)

	r.Use(authentication.Middleware)
}
