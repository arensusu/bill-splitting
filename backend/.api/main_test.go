package api

import (
	db "bill-splitting/db/sqlc"
	"bill-splitting/token"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func newTestServer(store db.Store) *Server {
	tokenMaker := token.NewJWTMaker("secret")
	server := NewServer(store, tokenMaker)
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
