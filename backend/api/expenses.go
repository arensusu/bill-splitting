package api

import (
	db "bill-splitting/db/sqlc"
	"bill-splitting/token"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type createExpenseUriRequest struct {
	GroupID int32 `uri:"groupId" binding:"required"`
}
type createExpenseJSONRequest struct {
	Amount      int64  `json:"amount" binding:"required"`
	Description string `json:"description"`
	Date        string `json:"date" binding:"required"`
}

func (s *Server) createExpense(ctx *gin.Context) {
	var uriRequest createExpenseUriRequest
	if err := ctx.ShouldBindUri(&uriRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var jsonRequest createExpenseJSONRequest
	if err := ctx.ShouldBindJSON(&jsonRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payload := ctx.MustGet("payload").(*token.JWTPayload)

	member, err := s.store.GetMembership(ctx, db.GetMembershipParams{
		GroupID: uriRequest.GroupID,
		UserID:  payload.UserID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if payload.UserID != member.UserID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized user"})
		return
	}

	date, err := time.Parse("2006-01-02", jsonRequest.Date)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expenseTx, err := s.store.CreateExpense(ctx, db.CreateExpenseParams{
		MemberID:    member.ID,
		Amount:      strconv.FormatInt(jsonRequest.Amount, 10),
		Description: jsonRequest.Description,
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
