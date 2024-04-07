package api

import (
	"bill-splitting/helper"
	"bill-splitting/token"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func addAuthentication(
	t *testing.T,
	request *http.Request,
	maker *token.JWTMaker,
	userID string,
	duration time.Duration,
) {
	token, payload, err := maker.CreateToken(userID, duration)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	request.Header.Set("Authorization", "Bearer "+token)
}

func TestAuthMiddleware(t *testing.T) {
	userID := helper.RandomString(10)
	duration := 15 * time.Minute

	server := newTestServer(nil)
	server.router.GET("/auth", authMiddleware(server.tokenMaker), func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{})
	})

	recorder := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/auth", nil)
	require.NoError(t, err)

	addAuthentication(t, request, server.tokenMaker, userID, duration)
	server.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)
}
