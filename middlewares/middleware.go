package middlewares

import "github.com/gorilla/mux"

const (
	UserContext = "UserContext"
)

func RegisterMiddlewares(r *mux.Router) {
	var (
		authentication = Authentication{WhitelistedRoutes: []string{"/login"}}
	)

	r.Use(authentication.Middleware)
}
