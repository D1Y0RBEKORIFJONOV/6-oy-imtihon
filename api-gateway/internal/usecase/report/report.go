package reportusecase

import (
	"apigateway/internal/entity"
	"context"
)

type reportUseCase interface {
	GetSpendingByCategory(ctx context.Context, id string) (*entity.ListSpendingResponse, error)
	GetIncomeExpense(ctx context.Context, id string) (*entity.IncomeExpenseResponse, error)
	GeFromTill(ctx context.Context, req *entity.FromTillRequest) (*entity.FromTillResponse, error)
}

type ReportUseCaseImpl struct {
	report reportUseCase
}

func NewReportUseCase(report reportUseCase) *ReportUseCaseImpl {
	return &ReportUseCaseImpl{
		report: report,
	}
}
func (r *ReportUseCaseImpl) GetSpendingByCategory(ctx context.Context, id string) (*entity.ListSpendingResponse, error) {
	return r.report.GetSpendingByCategory(ctx, id)
}
func (r *ReportUseCaseImpl) GetIncomeExpense(ctx context.Context, id string) (*entity.IncomeExpenseResponse, error) {
	return r.report.GetIncomeExpense(ctx, id)
}

func (r *ReportUseCaseImpl) GeFromTill(ctx context.Context, req *entity.FromTillRequest) (*entity.FromTillResponse, error) {
	return r.report.GeFromTill(ctx, req)
}
