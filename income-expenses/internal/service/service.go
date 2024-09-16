package service

import (
	"context"
	interface17 "incomeexpenses/internal/interface"
	"incomeexpenses/internal/protos/income"
)

type Service struct {
	income.UnimplementedIncomeExpensesServer
	S interface17.IncomeExpenses
}

func (u *Service) Expenses(ctx context.Context, req *income.CreateIncomeExpensesRequest) (*income.CreateIncomeExpensesResponse, error) {
	res, err := u.S.Expenses(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (u *Service) Income(ctx context.Context, req *income.CreateIncomeExpensesRequest) (*income.CreateIncomeExpensesResponse, error) {
	res, err := u.S.Income(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (u *Service) Info(ctx context.Context, req *income.GetInfoRequest) (*income.GetInfoResponse, error) {
	res, err := u.S.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
