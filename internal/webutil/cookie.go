package webutil

import (
	"net/http"
	"time"
)

const (
	SessionCookieName     = "session_id"
	SessionCookieDuration = 24 * time.Hour
	SessionCookieSecure   = false // true quando usar HTTPS
)

func SetSessionCookie(
	w http.ResponseWriter,
	sessionID string,
) {
	http.SetCookie(w, &http.Cookie{
		Name:     SessionCookieName,
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
		Secure:   SessionCookieSecure,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   int(SessionCookieDuration.Seconds()),
		Expires:  time.Now().Add(SessionCookieDuration),
	})
}

func ClearSessionCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     SessionCookieName,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})
}

func GetSessionCookie(r *http.Request) *string {
	cookie, err := r.Cookie(SessionCookieName)
	if err != nil {
		return nil
	}

	if cookie.Value == "" {
		return nil
	}

	return &cookie.Value
}
