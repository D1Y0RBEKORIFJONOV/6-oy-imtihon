package interface17

import (
	"context"
	"incomeexpenses/internal/protos/income"
)

type IncomeExpenses interface {
	Income(ctx context.Context, req *income.CreateIncomeExpensesRequest) (*income.CreateIncomeExpensesResponse, error)
	Expenses(ctx context.Context, req *income.CreateIncomeExpensesRequest) (*income.CreateIncomeExpensesResponse, error)
	Get(ctx context.Context, req *income.GetInfoRequest) (*income.GetInfoResponse, error)
}
