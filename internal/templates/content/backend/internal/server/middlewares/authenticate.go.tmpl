package middlewares

import (
	"net/http"
)

func (m *Middlewares) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isLoginRequest := r.Pattern == "GET /login"

		if !m.disableAuthentication {
			userID := m.sessionManager.GetString(r.Context(), "userID")
			validAuthToken := userID != ""

			if !validAuthToken && !isLoginRequest {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}
			if validAuthToken && isLoginRequest {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
		}

		if m.disableAuthentication && isLoginRequest {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}
