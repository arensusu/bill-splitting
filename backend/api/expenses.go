package api

import (
	db "bill-splitting/db/sqlc"
	"bill-splitting/token"
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/freetype/truetype"
	"github.com/wcharczuk/go-chart/v2"
)

type createExpenseUriRequest struct {
	GroupID int32 `uri:"groupId" binding:"required"`
}
type createExpenseJSONRequest struct {
	Category    string `json:"category"`
	Amount      string `json:"amount" binding:"required"`
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
		Amount:      jsonRequest.Amount,
		Description: jsonRequest.Description,
		Date:        date,
		Category:    sql.NullString{String: jsonRequest.Category, Valid: true},
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

func (s *Server) listExpensesSummary(c *gin.Context) {
	var req listExpensesRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	summary, err := s.store.ListSumOfExpensesWithCategory(c, req.GroupID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	data, err := json.Marshal(summary)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	hasher := sha256.New()
	hasher.Write(data)
	hashBytes := hasher.Sum(nil)

	values := make([]chart.Value, len(summary))
	for i, v := range summary {
		values[i] = chart.Value{
			Value: float64(v.Sum),
			Label: fmt.Sprintf("%s $%d", v.Category.String, v.Sum),
		}
	}

	fontBytes, err := os.ReadFile("./msjh.ttc")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	font, err := truetype.Parse(fontBytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pie := chart.PieChart{
		Width:  512,
		Height: 512,
		Values: values,
		Font:   font,
	}

	f, err := os.Create(fmt.Sprintf("/var/images/%x.png", hashBytes))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer f.Close()
	err = pie.Render(chart.PNG, f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"image": fmt.Sprintf("%x.png", hashBytes)})
}
