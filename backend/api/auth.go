package api

import (
	db "bill-splitting/db/sqlc"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func (s *Server) auth(ctx *gin.Context) {
	gothic.GetProviderName = func(r *http.Request) (string, error) { return ctx.Param("provider"), nil }

	res := ctx.Writer
	req := ctx.Request

	if _, err := gothic.CompleteUserAuth(res, req); err != nil {
		gothic.BeginAuthHandler(res, req)
	}
}

func (s *Server) authCallback(ctx *gin.Context) {
	gothic.GetProviderName = func(r *http.Request) (string, error) { return ctx.Param("provider"), nil }
	gotUser, err := gothic.CompleteUserAuth(ctx.Writer, ctx.Request)
	if err != nil {
		fmt.Fprintln(ctx.Writer, err)
		return
	}

	userID := gotUser.UserID
	username := gotUser.Name
	if username == "" {
		username = gotUser.NickName
	}

	user, err := s.store.GetUser(ctx, userID)
	if err != nil {
		if err != sql.ErrNoRows {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		user, err = s.store.CreateUser(ctx, db.CreateUserParams{
			ID:       userID,
			Username: username,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	token, _, err := s.tokenMaker.CreateToken(user.ID, time.Hour)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	endpoint := os.Getenv("ENDPOINT")
	ctx.SetCookie("token", token, int(time.Hour.Seconds()), "/", endpoint, false, true)
	ctx.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s/login", endpoint))
}

type authLineBotRequest struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func (s *Server) authLineBot(ctx *gin.Context) {
	var req authLineBotRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := s.store.GetUser(ctx, req.ID)
	if err != nil {
		if err != sql.ErrNoRows {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		user, err = s.store.CreateUser(ctx, db.CreateUserParams{
			ID:       req.ID,
			Username: req.Username,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	token, _, err := s.tokenMaker.CreateToken(user.ID, time.Hour)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
