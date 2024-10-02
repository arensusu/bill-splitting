package api

import (
	db "bill-splitting/db/sqlc"
	"bytes"
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

var lineOauthConfig = &oauth2.Config{
	ClientID:     os.Getenv("LINE_CLIENT_ID"),                                        // LINE Developers 取得
	ClientSecret: os.Getenv("LINE_CLIENT_SECRET"),                                    // LINE Developers 取得
	RedirectURL:  fmt.Sprintf("%s/api/v1/auth/line/callback", os.Getenv("ENDPOINT")), // 設定的回調 URL
	Scopes:       []string{"profile", "openid", "email"},
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://access.line.me/oauth2/v2.1/authorize",
		TokenURL: "https://api.line.me/oauth2/v2.1/token",
	},
}

var state string

func setState() string {
	// If a state query param is not passed in, generate a random
	// base64-encoded nonce so that the state on the auth URL
	// is unguessable, preventing CSRF attacks, as described in
	//
	// https://auth0.com/docs/protocols/oauth2/oauth-state#keep-reading
	nonceBytes := make([]byte, 64)
	_, err := io.ReadFull(rand.Reader, nonceBytes)
	if err != nil {
		panic("source of randomness unavailable: " + err.Error())
	}
	return base64.URLEncoding.EncodeToString(nonceBytes)
}

func (s *Server) auth(ctx *gin.Context) {
	state = setState()
	url := lineOauthConfig.AuthCodeURL(state)
	ctx.Redirect(http.StatusTemporaryRedirect, url)
}

func (s *Server) authCallback(ctx *gin.Context) {
	getState := ctx.Query("state")
	if state != getState {
		log.Printf("Invalid OAuth state")
		ctx.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	// 取得授權碼
	code := ctx.Query("code")
	lineToken, err := lineOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Printf("Failed to exchange token: %v", err)
		ctx.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	// 使用 token 獲取使用者資訊
	client := lineOauthConfig.Client(context.Background(), lineToken)
	resp, err := client.Get("https://api.line.me/v2/profile")
	if err != nil {
		log.Printf("Failed to get user info: %v", err)
		ctx.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		ctx.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	lineUser := struct {
		Name          string `json:"name"`
		UserID        string `json:"userId"`
		DisplayName   string `json:"displayName"`
		PictureURL    string `json:"pictureUrl"`
		StatusMessage string `json:"statusMessage"`
	}{}

	if err = json.NewDecoder(bytes.NewReader(body)).Decode(&lineUser); err != nil {
		log.Printf("Failed to decode user info: %v", err)
		ctx.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	userID := lineUser.UserID
	username := lineUser.Name

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

	//endpoint := os.Getenv("ENDPOINT")
	ctx.SetCookie("token", token, int(time.Hour.Seconds()), "/", "", false, true)
	ctx.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("/"))
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
