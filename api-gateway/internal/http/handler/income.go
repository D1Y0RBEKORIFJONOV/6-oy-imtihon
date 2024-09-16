package handler

import (
	"apigateway/internal/entity"
	incomeusecase "apigateway/internal/usecase/income"

	"github.com/gin-gonic/gin"
	"net/http"
)

type Income struct {
	income incomeusecase.IncomeUseCaseIml
}

func NewIncome(income incomeusecase.IncomeUseCaseIml) *Income {
	return &Income{
		income: income,
	}
}

// CreateIncome godoc
// @Summary CreateIncome
// @Description create a new income
// @Tags income
// @Accept json
// @Produce json
// @Param body body entity.CreateIncomeExpensesRequest true "Budget create information"
// @Security ApiKeyAuth
// @Success 200 {object} entity.CreateIncomeExpensesResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /income/create/income [post]
// @Security BearerAuth
func (i *Income) CreateIncome(c *gin.Context) {
	var req entity.CreateIncomeExpensesRequest
	id, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No UserID in request context"})
		return
	}
	idStr, ok := id.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No ID in request context"})
		return
	}
	req.UserID = idStr
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := i.income.CreateIncome(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetIncome godoc
// @Summary GetIncome
// @Description GetIncome  income
// @Tags income
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} entity.GetInfoResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /income [get]
// @Security BearerAuth
func (i *Income) GetIncome(c *gin.Context) {
	var req entity.GetInfoRequest
	id, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No UserID in request context"})
		return
	}
	idStr, ok := id.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No ID in request context"})
		return
	}
	req.UserID = idStr

	res, err := i.income.GetInfo(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// Expenses godoc
// @Summary Expenses
// @Description create a new income
// @Tags income
// @Accept json
// @Produce json
// @Param body body entity.CreateIncomeExpensesRequest true "Budget create information"
// @Security ApiKeyAuth
// @Success 200 {object} entity.CreateIncomeExpensesResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /income/create/expenses [post]
// @Security BearerAuth
func (i *Income) Expenses(c *gin.Context) {
	var req entity.CreateIncomeExpensesRequest
	id, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No UserID in request context"})
		return
	}
	idStr, ok := id.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No ID in request context"})
		return
	}
	req.UserID = idStr
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := i.income.CreateExpenses(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
