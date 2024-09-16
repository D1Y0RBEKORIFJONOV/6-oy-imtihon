package budgetservice

import (
	budgetproto "apigateway/gen/go/budget"
	"apigateway/internal/entity"
	clientgrpcserver "apigateway/internal/infastructure/client_grpc_server"
	"context"
	"fmt"
	"log/slog"
)

type Budget struct {
	client clientgrpcserver.ServiceClient
	logger *slog.Logger
}

func NewBudget(client clientgrpcserver.ServiceClient, logger *slog.Logger) *Budget {
	return &Budget{

		client: client,
		logger: logger,
	}
}

func (b *Budget) CreateBudget(ctx context.Context, req *entity.CreateBudgetRequest) (*entity.CreateBudgetResponse, error) {
	const op = "Service.CreateBudget"
	log := b.logger.With(
		slog.String("method", op))
	log.Info("Start")
	defer log.Info("End")

	budget, err := b.client.BudgetServiceClient().CreateBudget(ctx, &budgetproto.CreateBudgetRequest{
		UserId:   req.UserID,
		Category: req.Category,
		Amount:   req.Amount,
		Currency: req.Currency,
	})
	if err != nil {
		log.Error("err", err)
		return nil, err
	}

	return &entity.CreateBudgetResponse{
		Message:  budget.Message,
		BudgetID: budget.BudgetId,
	}, nil
}

func (b *Budget) GetBudgets(ctx context.Context, req *entity.GetBudgetsRequest) (*entity.GetBudgetsResponse, error) {
	const op = "Service.GetBudgets"
	log := b.logger.With(
		slog.String("method", op))
	log.Info("Start")
	defer log.Info("End")

	fmt.Printf("%+v", req)
	budgets, err := b.client.BudgetServiceClient().GetBudgets(ctx, &budgetproto.GetBudgetsRequest{
		UserId:   req.UserID,
		Category: req.Category,
	})
	if err != nil {
		log.Error("err", err)
		return nil, err
	}
	budget := entity.Budget{
		BudgetID: budgets.Budgets.BudgetId,
		Category: budgets.Budgets.Category,
		Amount:   budgets.Budgets.Amount,
		Spent:    budgets.Budgets.Spent,
		Currency: budgets.Budgets.Currency,
		UserID:   budgets.Budgets.UserId}

	return &entity.GetBudgetsResponse{
		Budgets: budget,
	}, nil
}

func (b *Budget) UpdateBudget(ctx context.Context, req *entity.UpdateBudgetRequest) (*entity.UpdateBudgetResponse, error) {
	const op = "Service.UpdateBudget"
	log := b.logger.With(
		slog.String("method", op))
	res, err := b.client.BudgetServiceClient().UpdateBudget(ctx, &budgetproto.UpdateBudgetRequest{
		UserId:   req.UserID,
		Amount:   req.Amount,
		Spent:    req.Spent,
		Currency: req.Currency,
		BudgetId: req.BudgetID,
	})
	if err != nil {
		log.Error("err", err)
		return nil, err
	}

	return &entity.UpdateBudgetResponse{
		Message: res.Message,
	}, nil
}
