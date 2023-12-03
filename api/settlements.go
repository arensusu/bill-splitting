package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type replaceSettlementRequest struct {
	GroupID int64 `json:"groupId" binding:"required"`
}

func (s *Server) replaceSettlement(c *gin.Context) {
	var req replaceSettlementRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := s.store.CreateSettlementsTx(c, req.GroupID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
