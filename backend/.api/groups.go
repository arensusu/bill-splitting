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
	group, err := s.store.CreateGroupTx(c, db.CreateGroupTxParams{
		Name:   req.Name,
		UserID: payload.UserID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, group)
}

type getGroupRequest struct {
	ID int32 `uri:"groupId" binding:"required,min=1"`
}

func (s *Server) getGroup(c *gin.Context) {
	var req getGroupRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	group, err := s.store.GetGroup(c, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, group)
}

func (s *Server) listGroups(c *gin.Context) {
	payload := c.MustGet("payload").(*token.JWTPayload)

	groups, err := s.store.ListGroups(c, payload.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, groups)
}
