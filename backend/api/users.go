package api

// import (
// 	db "bill-splitting/db/sqlc"
// 	"database/sql"
// 	"fmt"
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"golang.org/x/crypto/bcrypt"
// )

// type createUserRequest struct {
// 	Username string `json:"username" binding:"required"`
// 	Password string `json:"password" binding:"required"`
// }

// type createUserResponse struct {
// 	ID        string    `json:"id"`
// 	Username  string    `json:"username"`
// 	CreatedAt time.Time `json:"createdAt"`
// }

// func (s *Server) createUser(c *gin.Context) {
// 	var req createUserRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	user, err := s.store.CreateUser(c, db.CreateUserParams{
// 		Username: req.Username,
// 	})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, createUserResponse{
// 		ID:        user.ID,
// 		Username:  user.Username,
// 		CreatedAt: user.CreatedAt,
// 	})
// }

// type loginUserRequest struct {
// 	Username string `json:"username" binding:"required"`
// 	Password string `json:"password" binding:"required"`
// }

// type loginUserResponse struct {
// 	Token string             `json:"token"`
// 	Exp   time.Time          `json:"exp"`
// 	User  createUserResponse `json:"user"`
// }

// func (s *Server) loginUser(c *gin.Context) {
// 	var req loginUserRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	user, err := s.store.GetUserByUsername(c, req.Username)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
// 			return
// 		}
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
// 		return
// 	}

// 	token, payload, err := s.tokenMaker.CreateToken(user.ID, time.Hour)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, loginUserResponse{
// 		Token: token,
// 		Exp:   payload.ExpiresAt.Time,
// 		User: createUserResponse{
// 			ID:        user.ID,
// 			Username:  user.Username,
// 			CreatedAt: user.CreatedAt,
// 		},
// 	})
// }
