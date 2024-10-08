package api

import (
	db "bill-splitting/db/sqlc"
	"bill-splitting/token"
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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

type listExpensesSummaryQueryParams struct {
	StartTime string `form:"startTime"`
	EndTime   string `form:"endTime"`
}

// func (s *Server) listExpensesSummary(c *gin.Context) {
// 	var req listExpensesRequest
// 	if err := c.ShouldBindUri(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	var query listExpensesSummaryQueryParams
// 	if err := c.ShouldBindQuery(&query); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	startTime, err := time.Parse("2006-01-02", query.StartTime)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	endTime, err := time.Parse("2006-01-02", query.EndTime)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	summary, err := s.store.SummarizeExpensesWithinDate(c, db.SummarizeExpensesWithinDateParams{
// 		GroupID:   req.GroupID,
// 		StartTime: startTime,
// 		EndTime:   endTime,
// 	})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	data, err := json.Marshal(summary)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	hasher := sha256.New()
// 	hasher.Write(data)
// 	hashBytes := hasher.Sum(nil)

// 	values := make([]chart.Value, len(summary))
// 	for i, v := range summary {
// 		values[i] = chart.Value{
// 			Value: float64(v.Total),
// 			Label: fmt.Sprintf("%s $%d", v.Category.String, v.Total),
// 		}
// 	}

// 	fontBytes, err := os.ReadFile("./msjh.ttc")
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	font, err := truetype.Parse(fontBytes)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	pie := chart.PieChart{
// 		Title:  fmt.Sprintf("%s ~ %s", query.StartTime, query.EndTime),
// 		Width:  500,
// 		Height: 600,
// 		Values: values,
// 		Font:   font,
// 	}

// 	f, err := os.Create(fmt.Sprintf("/var/images/%x.png", hashBytes))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	defer f.Close()
// 	err = pie.Render(chart.PNG, f)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"image": fmt.Sprintf("%x.png", hashBytes)})
// }
