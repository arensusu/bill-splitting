package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type createSettlementRequest struct {
	GroupID int64 `json:"groupId" binding:"required"`
}

func (s *Server) createSettlement(c *gin.Context) {
	var req createSettlementRequest
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
