package incomeservice

import (
	incomepb "apigateway/gen/go/incomeexpences"
	"apigateway/internal/config"
	"apigateway/internal/entity"
	clientgrpcserver "apigateway/internal/infastructure/client_grpc_server"
	"apigateway/internal/infastructure/producer"
	redisrepository "apigateway/internal/infastructure/repository/redis"
	"context"
	"encoding/json"
	"log/slog"
)

type Income struct {
	redis  *redisrepository.RedisUserRepository
	logger *slog.Logger
	client clientgrpcserver.ServiceClient
	cfg    *config.Config
}

func NewIncome(logger *slog.Logger, client clientgrpcserver.ServiceClient, redis *redisrepository.RedisUserRepository) *Income {
	cfg := config.New()
	return &Income{
		logger: logger,
		client: client,
		redis:  redis,
		cfg:    cfg,
	}
}

func (i *Income) CreateIncome(ctx context.Context, req *entity.CreateIncomeExpensesRequest) (resp *entity.CreateIncomeExpensesResponse, err error) {
	const op = "Service.CreateIncome"
	log := i.logger.With(
		slog.String("method", op))
	log.Info("start")
	defer log.Info("end")

	err = i.createValue(ctx, req, i.cfg.MessageBrokerUses.TopicIncome, i.cfg.MessageBrokerUses.IncomeCreate)
	if err != nil {
		return nil, err
	}
	err = i.redis.DeleteFromCache(ctx, "budget:"+req.UserID+":"+req.Category)
	if err != nil {
		log.Error("err", err)
		return nil, err
	}
	return &entity.CreateIncomeExpensesResponse{
		Message: "creating income in progress",
	}, nil
}

func (i *Income) CreateExpenses(ctx context.Context, req *entity.CreateIncomeExpensesRequest) (resp *entity.CreateIncomeExpensesResponse, err error) {
	const op = "Service.Expenses"
	log := i.logger.With(
		slog.String("method", op))
	log.Info("start")
	defer log.Info("end")

	err = i.createValue(ctx, req, i.cfg.MessageBrokerUses.TopicIncome, i.cfg.MessageBrokerUses.ExpensesCreate)
	if err != nil {
		return nil, err
	}
	err = i.redis.DeleteFromCache(ctx, "budget:"+req.UserID+":"+req.Category)
	if err != nil {
		log.Error("err", err)
		return nil, err
	}
	return &entity.CreateIncomeExpensesResponse{
		Message: "creating expenses in progress",
	}, nil
}

func (i *Income) createValue(ctx context.Context, req *entity.CreateIncomeExpensesRequest, topicName, key string) error {
	reqProto := incomepb.CreateIncomeExpensesRequest{
		UserId:   req.UserID,
		Category: req.Category,
		Amount:   req.Amount,
		Date:     req.Date,
		Type:     req.Type,
		Currency: req.Currency,
	}
	reqByte, err := json.Marshal(&reqProto)
	if err != nil {
		return err
	}
	reqKey := []byte(key)
	err = producer.Producer(ctx, reqKey, reqByte, topicName)
	if err != nil {
		return err
	}

	return nil
}

func (i *Income) GetInfo(ctx context.Context, req *entity.GetInfoRequest) (*entity.GetInfoResponse, error) {
	const op = "Service.GetInfo"
	log := i.logger.With(
		slog.String("method", op))

	log.Info("start")
	defer log.Info("end")

	res, err := i.client.IncomeServiceClient().Info(ctx, &incomepb.GetInfoRequest{
		UserId: req.UserID,
	})
	if err != nil {
		log.Error("err", err)
		return nil, err
	}
	var response []entity.InfoResponse
	for _, info := range res.Info {
		response = append(response, entity.InfoResponse{
			TransactionID: info.Transactionid,
			Type:          info.Type,
			Category:      info.Category,
			Currency:      info.Currency,
			Amount:        info.Amount,
			Date:          info.Date,
			UserID:        info.UserId,
		})
	}

	return &entity.GetInfoResponse{
		Info: response,
	}, nil
}
