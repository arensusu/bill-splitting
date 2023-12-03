package api

import (
	db "bill-splitting/db/sqlc"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type createExpenseRequest struct {
	GroupID     int64     `json:"groupId" binding:"required"`
	PayerID     int64     `json:"payerId" binding:"required"`
	Amount      int64     `json:"amount" binding:"required"`
	Description string    `json:"description"`
	Date        time.Time `json:"date" binding:"required"`
}

func (s *Server) createExpense(c *gin.Context) {
	var req createExpenseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := s.store.GetGroupMember(c, db.GetGroupMemberParams{
		GroupID: req.GroupID,
		UserID:  req.PayerID,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusForbidden, gin.H{"error": errors.New("user is not a member of the group")})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	expenseTx, err := s.store.CreateExpenseTx(c, db.CreateExpenseTxParams{
		GroupID:     req.GroupID,
		PayerID:     req.PayerID,
		Amount:      req.Amount,
		Description: req.Description,
		Date:        req.Date,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, expenseTx)
}

type listExpensesRequest struct {
	GroupID int64 `uri:"groupId" binding:"required"`
}

func (s *Server) listExpenses(c *gin.Context) {
	var req listExpensesRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expenses, err := s.store.ListExpenses(c, req.GroupID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, expenses)
}
