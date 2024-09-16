package budgetusecase

import (
	"apigateway/internal/entity"
	"context"
)

type budgetUseCase interface {
	CreateBudget(ctx context.Context, req *entity.CreateBudgetRequest) (*entity.CreateBudgetResponse, error)
	GetBudgets(ctx context.Context, req *entity.GetBudgetsRequest) (*entity.GetBudgetsResponse, error)
	UpdateBudget(ctx context.Context, req *entity.UpdateBudgetRequest) (*entity.UpdateBudgetResponse, error)
}

type BudgetUseCaseImpl struct {
	budget budgetUseCase
}

func NewBudgetUseCase(budget budgetUseCase) *BudgetUseCaseImpl {
	return &BudgetUseCaseImpl{budget: budget}
}
func (b *BudgetUseCaseImpl) CreateBudget(ctx context.Context, req *entity.CreateBudgetRequest) (*entity.CreateBudgetResponse, error) {
	return b.budget.CreateBudget(ctx, req)
}
func (b *BudgetUseCaseImpl) GetBudgets(ctx context.Context, req *entity.GetBudgetsRequest) (*entity.GetBudgetsResponse, error) {
	return b.budget.GetBudgets(ctx, req)
}

func (b *BudgetUseCaseImpl) UpdateBudget(ctx context.Context, req *entity.UpdateBudgetRequest) (*entity.UpdateBudgetResponse, error) {
	return b.budget.UpdateBudget(ctx, req)
}
