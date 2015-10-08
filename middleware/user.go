package middleware

import (
	"net/http"

	"ojapi/config"
	"ojapi/session"

	"github.com/zenazn/goji/web"
)

func SetUser(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var user = session.GetUser(r)
		if user != nil {
			UserToC(c, user)
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// RequireUserAdmin is a middleware function that verifies
// there is a currently authenticated user stored in
// the context with ADMIN privilege.
func RequireUserAdmin(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var user = ToUser(*c)
		switch {
		case user == nil:
			w.WriteHeader(http.StatusUnauthorized)
			return
		case user != nil && (user.Privilege != config.PrivilegeAD && user.Privilege != config.PrivilegeTC):
			w.WriteHeader(http.StatusForbidden)
			return
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
