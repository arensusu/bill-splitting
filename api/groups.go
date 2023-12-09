package api

import (
	db "bill-splitting/db/sqlc"
	"bill-splitting/token"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createGroupRequest struct {
	Name string `json:"name" binding:"required"`
}

func (s *Server) createGroup(c *gin.Context) {
	var req createGroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payload := c.MustGet("payload").(*token.JWTPayload)
	user, err := s.store.CreateGroupTx(c, db.CreateGroupTxParams{
		Name:   req.Name,
		UserID: payload.UserID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

type getGroupRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (s *Server) getGroup(c *gin.Context) {
	var req getGroupRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := s.store.GetGroup(c, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
