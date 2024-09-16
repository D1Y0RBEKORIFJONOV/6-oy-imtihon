package handler

import (
	"apigateway/internal/entity"
	budgetusecase "apigateway/internal/usecase/budget"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BudgetHandler struct {
	budget *budgetusecase.BudgetUseCaseImpl
}

func NewBudgetHandler(budget *budgetusecase.BudgetUseCaseImpl) *BudgetHandler {
	return &BudgetHandler{budget: budget}
}

// CreateBudget godoc
// @Summary CreateBudget
// @Description create a new budget
// @Tags budget
// @Accept json
// @Produce json
// @Param body body entity.CreateBudgetRequest true "Budget create information"
// @Security ApiKeyAuth
// @Success 200 {object} entity.CreateBudgetResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /budget/create [post]
// @Security BearerAuth
func (h *BudgetHandler) CreateBudget(c *gin.Context) {
	var req entity.CreateBudgetRequest

	id, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id not found"})
		return
	}
	idString, ok := id.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id not found"})
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.UserID = idString

	res, err := h.budget.CreateBudget(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetBudget godoc
// @Summary GetBudget
// @Description get a user budget
// @Tags budget
// @Accept json
// @Produce json
// @Param category path string true "category"
// @Security ApiKeyAuth
// @Success 200 {object} entity.GetBudgetsResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /budget/{category} [get]
// @Security BearerAuth
func (h *BudgetHandler) GetBudget(c *gin.Context) {
	id, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id not found"})
		return
	}
	idString, ok := id.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id not found"})
	}
	category := c.Param("category")
	if category == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "category not found"})
		return
	}
	var req entity.GetBudgetsRequest
	req.UserID = idString
	req.Category = category

	res, err := h.budget.GetBudgets(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// UpdateBudget godoc
// @Summary UpdateBudget
// @Description update a user budget
// @Tags budget
// @Accept json
// @Produce json
// @Param body body entity.UpdateBudgetRequest true "Budget create information"
// @Param budget_id path string true "budget_id"
// @Security ApiKeyAuth
// @Success 200 {object} entity.UpdateBudgetResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /budget/{budget_id} [patch]
// @Security BearerAuth
func (h *BudgetHandler) UpdateBudget(c *gin.Context) {
	var req entity.UpdateBudgetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id not found"})
		return
	}
	idString, ok := id.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id not found"})
		return
	}
	budgetID := c.Param("budget_id")

	if budgetID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "budget_id not found "})
		return
	}

	req.BudgetID = budgetID
	req.UserID = idString
	res, err := h.budget.UpdateBudget(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
