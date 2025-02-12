package identity

import (
	"context"
	"net/http"
)

type userkey int

const (
	_ userkey = iota
	key
)

func ContextWithUser(ctx context.Context, user string) context.Context {
	return context.WithValue(ctx, key, user)
}

func UserFromContext(ctx context.Context) (string, bool) {
	user, ok := ctx.Value(key).(string)
	return user, ok
}

func extractUser(r *http.Request) (string, error) {
	userCookie, err := r.Cookie("identity")
	if err != nil {
		return "", err
	}
	return userCookie.Value, nil
}

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := extractUser(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("unauthorized"))
			return
		}
		ctx := r.Context()
		ctx = ContextWithUser(ctx, user)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}

func setUser(user string, w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:  "identity",
		Value: user,
	})
}

func DeleteUser(rw http.ResponseWriter) {
	http.SetCookie(rw, &http.Cookie{
		Name:   "identity",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
}