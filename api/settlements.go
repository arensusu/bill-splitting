package api

import (
	db "bill-splitting/db/sqlc"
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

type completeSettlementRequest struct {
	GroupID int64  `json:"groupId" binding:"required"`
	PayerID string `json:"payerId" binding:"required"`
	PayeeID string `json:"payeeId" binding:"required"`
}

func (s *Server) completeSettlement(c *gin.Context) {
	var req completeSettlementRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := s.store.DeleteSettlement(c, db.DeleteSettlementParams{
		GroupID: req.GroupID,
		PayerID: req.PayerID,
		PayeeID: req.PayeeID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}
