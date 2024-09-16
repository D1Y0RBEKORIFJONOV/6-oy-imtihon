package reportservice

import (
	reportproto "apigateway/gen/go/report"
	"apigateway/internal/entity"
	clientgrpcserver "apigateway/internal/infastructure/client_grpc_server"
	"context"
	"log/slog"
)

type Report struct {
	logger *slog.Logger
	client clientgrpcserver.ServiceClient
}

func NewReport(logger *slog.Logger, client clientgrpcserver.ServiceClient) *Report {
	return &Report{logger: logger, client: client}
}

func (r *Report) GetSpendingByCategory(ctx context.Context, id string) (*entity.ListSpendingResponse, error) {
	const op = "Service.GetSpendingByCategory"
	log := r.logger.With(
		"method", r.logger)
	log.Info("start")
	defer log.Info("end")
	res, err := r.client.ReportServiceClient().SpendingbyCategory(ctx, &reportproto.SpendingRequest{
		UserId: id,
	})
	if err != nil {
		log.Error(op, err)
		return nil, err
	}
	var response []entity.SpendingResponse
	for _, spending := range res.Spent {
		response = append(response, entity.SpendingResponse{
			Category:   spending.Category,
			TotalSpent: spending.Totalspent,
		})
	}
	return &entity.ListSpendingResponse{
		Spent: response,
	}, nil
}

func (r *Report) GetIncomeExpense(ctx context.Context, id string) (*entity.IncomeExpenseResponse, error) {
	const op = "Service.GetIncomeExpense"
	log := r.logger.With(
		"method", r.logger)
	log.Info("start")
	defer log.Info("end")

	res, err := r.client.ReportServiceClient().IncomeExpense(ctx, &reportproto.SpendingRequest{
		UserId: id,
	})
	if err != nil {
		log.Error(op, err)
		return nil, err
	}

	return &entity.IncomeExpenseResponse{
		TotalIncome:   res.TotalIncome,
		TotalExpenses: res.TotalExpenses,
		NetSavings:    res.NetSavings,
	}, nil
}

func (r *Report) GeFromTill(ctx context.Context, req *entity.FromTillRequest) (*entity.FromTillResponse, error) {
	const op = "Service.GeFromTill"
	log := r.logger.With(
		"method", r.logger)
	log.Info("start")
	defer log.Info("end")
	res, err := r.client.ReportServiceClient().FromTill(ctx, &reportproto.FromTillRequest{
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		UserId:    req.UserID,
	})
	if err != nil {
		log.Error(op, err)
		return nil, err
	}

	return &entity.FromTillResponse{
		TotalAmount: res.TotalAmount,
	}, nil
}
