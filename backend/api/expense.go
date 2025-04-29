package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type listExpensesRequest struct {
	StartDate string `form:"start_date"`
	EndDate   string `form:"end_date"`
}

func (s *Server) listExpenses(c *gin.Context) {
	var req listExpensesRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	discordChannel := c.Param("discordChannel")

	group, err := s.store.GetGroupByDiscordChannel(discordChannel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	startTime, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	endTime, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expenses, err := s.store.ListExpensesWithinDate(group.ID, startTime, endTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, expenses)
}
