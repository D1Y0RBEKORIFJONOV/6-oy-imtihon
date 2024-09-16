package handler

import (
	"apigateway/internal/entity"
	reportusecase "apigateway/internal/usecase/report"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Report struct {
	report reportusecase.ReportUseCaseImpl
}

func NewReport(report reportusecase.ReportUseCaseImpl) *Report {
	return &Report{
		report: report,
	}
}

// GetSpending godoc
// @Summary GetSpending
// @Description GetSpending  speding
// @Tags report
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} entity.ListSpendingResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /report/spending [get]
// @Security BearerAuth
func (r *Report) GetSpending(c *gin.Context) {

	id, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No UserID in request context"})
		return
	}

	idStr, ok := id.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No UserID in request context"})
		return
	}

	res, err := r.report.GetSpendingByCategory(c.Request.Context(), idStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetExpenses godoc
// @Summary GetExpenses
// @Description GetExpenses  expenses
// @Tags report
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} entity.IncomeExpenseResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /report/expenses [get]
// @Security BearerAuth
func (r *Report) GetExpenses(c *gin.Context) {

	id, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No UserID in request context"})
		return
	}
	idStr, ok := id.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No UserID in request context"})
		return
	}

	res, err := r.report.GetIncomeExpense(c.Request.Context(), idStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetFromTill godoc
// @Summary GetFromTill
// @Description GetFromTill  expenses
// @Tags report
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param start_at path string true "start_at"
// @Param end_at path string true "end_at"
// @Success 200 {object} entity.FromTillResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /report/chosen/{start_at}/{end_at} [get]
// @Security BearerAuth
func (r *Report) GetFromTill(c *gin.Context) {
	id, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No UserID in request context"})
		return
	}
	idStr, ok := id.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No UserID in request context"})
		return
	}
	var req entity.FromTillRequest
	startAt := c.Param("start_at")
	endAt := c.Param("end_at")
	if startAt == "" || endAt == "" {
		log.Fatal(startAt, endAt)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter"})
		return
	}
	req.StartTime = startAt
	req.EndTime = endAt
	req.UserID = idStr

	res, err := r.report.GeFromTill(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
