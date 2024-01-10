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

// type loginUserResponse struct {
// 	Token string    `json:"token"`
// 	Exp   time.Time `json:"exp"`
// 	User  db.User   `json:"user"`
// }

func (s *Server) auth(ctx *gin.Context) {
	gothic.GetProviderName = func(r *http.Request) (string, error) { return ctx.Param("provider"), nil }

	res := ctx.Writer
	req := ctx.Request

	if gothUser, err := gothic.CompleteUserAuth(res, req); err == nil {
		fmt.Println(gothUser)
	} else {
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

	user, err := s.store.GetUser(ctx, gotUser.UserID)
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

	token, _, err := s.tokenMaker.CreateToken(user.ID, time.Hour)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	endpoint := os.Getenv("ENDPOINT")
	ctx.SetCookie("token", token, int(time.Hour.Seconds()), "/", endpoint, false, true)
	ctx.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s/login", endpoint))
}
