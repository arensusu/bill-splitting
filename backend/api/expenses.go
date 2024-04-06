package api

import (
	db "bill-splitting/db/sqlc"
	"bill-splitting/token"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type createExpenseRequest struct {
	MemberID    int32  `json:"memberId" binding:"required"`
	Amount      int64  `json:"amount" binding:"required"`
	Description string `json:"description"`
	Date        string `json:"date" binding:"required"`
}

func (s *Server) createExpense(ctx *gin.Context) {
	var req createExpenseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payload := ctx.MustGet("payload").(*token.JWTPayload)

	member, err := s.store.GetMember(ctx, req.MemberID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if payload.UserID != member.UserID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized user"})
		return
	}

	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expenseTx, err := s.store.CreateExpense(ctx, db.CreateExpenseParams{
		MemberID:    req.MemberID,
		Amount:      strconv.FormatInt(req.Amount, 10),
		Description: req.Description,
		Date:        date,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, expenseTx)
}

type listExpensesRequest struct {
	GroupID int32 `uri:"groupId" binding:"required"`
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
