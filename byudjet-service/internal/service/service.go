package service

import (
	"budgetservice/internal/entity/budget"
	"budgetservice/internal/infrastructura/repository"
	"context"
)

type BudgetService struct {
	repo  repository.BudgetRepository
}

func NewBudgetService(repo repository.BudgetRepository) *BudgetService {
	return &BudgetService{repo: repo}
}

func (b *BudgetService) Createbyudjet(ctx context.Context, req budget.CreateBudgetRequest) (*budget.CreateBudgetResponse, error) {
	return b.repo.AddMongo(ctx, req)
}

func (b *BudgetService) Getbyudjet(ctx context.Context, userID, category string) (*budget.Budget, error) {
	return b.repo.GetMongo(ctx, userID, category)
}

func (b *BudgetService) Updatebudget(ctx context.Context, req budget.UpdateBudgetRequest)(*budget.UpdateBudgetResponse, error){
	return b.repo.UpdateMongo(ctx, req)
}

func (b *BudgetService) Getusers(ctx context.Context, req budget.GetUserCategoriesRequest)(*budget.GetBudgetsResponse, error){
	return b.repo.GetUserMongo(ctx, req)
}
