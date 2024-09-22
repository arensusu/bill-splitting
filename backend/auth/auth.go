package auth

import (
	"fmt"
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/line"
)

const (
	key    = "secret"
	maxAge = 86400 * 30
	isProd = false
)

func NewAuth() {

	lineClientID := os.Getenv("LINE_CLIENT_ID")
	lineClientSecret := os.Getenv("LINE_CLIENT_SECRET")
	endpoint := os.Getenv("ENDPOINT")

	store := sessions.NewCookieStore([]byte(key))
	store.Options.Path = "/"
	store.Options.Domain = "localhost:8080"
	store.Options.MaxAge = maxAge
	store.Options.Secure = isProd

	gothic.Store = store

	goth.UseProviders(line.New(lineClientID, lineClientSecret, fmt.Sprintf("%s/api/v1/auth/line/callback", endpoint), "profile"))
}
