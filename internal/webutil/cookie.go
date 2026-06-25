package webutil

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

const (
	SessionCookieName     = "session_id"
	SessionCookieDuration = 24 * time.Hour
	SessionCookieSecure   = false // true quando usar HTTPS
)

func SetSessionCookie(
	w http.ResponseWriter,
	sessionID uuid.UUID,
) {
	http.SetCookie(w, &http.Cookie{
		Name:     SessionCookieName,
		Value:    sessionID.String(),
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

func GetSessionCookie(r *http.Request) *uuid.UUID {
	cookie, err := r.Cookie(SessionCookieName)
	if err != nil {
		return nil
	}

	if cookie.Value == "" {
		return nil
	}

	sessionID, err := uuid.Parse(cookie.Value)
	if err != nil {
		return nil
	}

	return &sessionID
}
