package incomeusecase

import (
	"apigateway/internal/entity"
	"context"
)

type incomeUseCase interface {
	CreateIncome(ctx context.Context, req *entity.CreateIncomeExpensesRequest) (resp *entity.CreateIncomeExpensesResponse, err error)
	CreateExpenses(ctx context.Context, req *entity.CreateIncomeExpensesRequest) (resp *entity.CreateIncomeExpensesResponse, err error)
	GetInfo(ctx context.Context, req *entity.GetInfoRequest) (resp *entity.GetInfoResponse, err error)
}

type IncomeUseCaseIml struct {
	income incomeUseCase
}

func NewIncomeUseCase(income incomeUseCase) *IncomeUseCaseIml {
	return &IncomeUseCaseIml{
		income: income,
	}
}

func (i *IncomeUseCaseIml) CreateIncome(ctx context.Context, req *entity.CreateIncomeExpensesRequest) (resp *entity.CreateIncomeExpensesResponse, err error) {
	return i.income.CreateIncome(ctx, req)
}

func (i *IncomeUseCaseIml) CreateExpenses(ctx context.Context, req *entity.CreateIncomeExpensesRequest) (resp *entity.CreateIncomeExpensesResponse, err error) {
	return i.income.CreateExpenses(ctx, req)
}

func (i *IncomeUseCaseIml) GetInfo(ctx context.Context, req *entity.GetInfoRequest) (resp *entity.GetInfoResponse, err error) {
	return i.income.GetInfo(ctx, req)
}
