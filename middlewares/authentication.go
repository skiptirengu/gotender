package middlewares

import (
	"github.com/skiptirengu/gotender/models"
	"net/http"
	"regexp"
	"github.com/skiptirengu/gotender/util"
	"github.com/gorilla/context"
)

const (
	header      = "Authorization"
	bearerRegex = "^Bearer\\s+(.*?)$"
)

type Authentication struct {
	WhitelistedRoutes []string
}

func (m Authentication) Middleware(next http.Handler) (http.Handler) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if m.isRouteWhitelisted(r.RequestURI) {
			next.ServeHTTP(w, r)
			return
		}
		requestToken := extractTokenFromHeader(r.Header.Get(header))
		if tokenModel := m.getToken(requestToken); tokenModel != nil {
			context.Set(r, UserContext, &tokenModel.User)
			next.ServeHTTP(w, r)
		} else {
			util.NewHttpError(w, http.StatusForbidden)
		}
	})
}

func (m Authentication) isRouteWhitelisted(route string) (bool) {
	for i := range m.WhitelistedRoutes {
		if m.WhitelistedRoutes[i] == route {
			return true
		}
	}
	return false
}

func (m Authentication) getToken(token string) (*models.Token) {
	if token == "" {
		return nil
	} else {
		return models.FindToken(token)
	}
}

func extractTokenFromHeader(header string) (string) {
	regex, _ := regexp.Compile(bearerRegex)
	matches := regex.FindStringSubmatch(header)
	if len(matches) == 2 {
		return matches[1]
	} else {
		return ""
	}
}
