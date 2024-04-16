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

	var user db.User
	userID := ctx.GetHeader("X-User")
	if userID != "" {
		var err error
		user, err = s.store.GetUser(ctx, userID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
	} else {
		gotUser, err := gothic.CompleteUserAuth(ctx.Writer, ctx.Request)
		if err != nil {
			fmt.Fprintln(ctx.Writer, err)
			return
		}

		user, err = s.store.GetUser(ctx, gotUser.UserID)
		if err != nil {
			if err != sql.ErrNoRows {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			username := gotUser.Name
			if username == "" {
				username = gotUser.NickName
			}

			user, err = s.store.CreateUser(ctx, db.CreateUserParams{
				ID:       gotUser.UserID,
				Username: username,
			})
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
	}

	token, _, err := s.tokenMaker.CreateToken(user.ID, time.Hour)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user.ID != "" {
		ctx.JSON(http.StatusOK, gin.H{"token": token})
	}

	endpoint := os.Getenv("ENDPOINT")
	ctx.SetCookie("token", token, int(time.Hour.Seconds()), "/", endpoint, false, true)
	ctx.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s/login", endpoint))
}
