package api

import (
	db "bill-splitting/db/sqlc"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func newTestServer(store db.Store) *Server {
	secretKey := "secret"
	server := NewServer(store, secretKey)
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
