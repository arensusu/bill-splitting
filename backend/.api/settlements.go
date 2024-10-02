package api

import (
	db "bill-splitting/db/sqlc"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type replaceSettlementRequest struct {
	GroupID int32 `uri:"groupId" binding:"required"`
}

func (s *Server) replaceSettlement(c *gin.Context) {
	var req replaceSettlementRequest
	if err := c.ShouldBindUri(&req); err != nil {
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
	PayerID int32 `uri:"payerId" binding:"required"`
	PayeeID int32 `uri:"payeeId" binding:"required"`
}

func (s *Server) completeSettlement(c *gin.Context) {
	var req completeSettlementRequest
	if err := c.ShouldBindUri(&req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := s.store.UpdateSettlement(c, db.UpdateSettlementParams{
		PayerID: req.PayerID,
		PayeeID: req.PayeeID,
		Amount:  "0",
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}
