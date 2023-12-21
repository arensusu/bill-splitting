package auth

import (
	"os"

	"github.com/gorilla/sessions"
	_ "github.com/joho/godotenv/autoload"
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

	store := sessions.NewCookieStore([]byte(key))
	store.Options.Path = "/"
	store.Options.MaxAge = maxAge
	store.Options.Secure = isProd

	gothic.Store = store

	goth.UseProviders(line.New(lineClientID, lineClientSecret, "http://localhost:8080/auth/line/callback", "profile"))
}
