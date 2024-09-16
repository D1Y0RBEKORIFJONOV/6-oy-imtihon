package repository

import (
	"budgetservice/internal/entity/budget"
	"context"
)

type BudgetRepository interface {
	AddMongo(ctx context.Context, req budget.CreateBudgetRequest) (*budget.CreateBudgetResponse, error)
	GetMongo(ctx context.Context, userID, category string) (*budget.Budget, error)
	UpdateMongo(ctx context.Context, req budget.UpdateBudgetRequest) (*budget.UpdateBudgetResponse, error)
	GetUserMongo(ctx context.Context, req budget.GetUserCategoriesRequest) (*budget.GetBudgetsResponse, error)
}
