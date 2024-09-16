package app

import (
	htppapp "apigateway/internal/app/htpp"
	"apigateway/internal/config"
	clientgrpcserver "apigateway/internal/infastructure/client_grpc_server"
	redisrepository "apigateway/internal/infastructure/repository/redis"
	budgetservice "apigateway/internal/service/budget"
	incomeservice "apigateway/internal/service/income"
	reportservice "apigateway/internal/service/report"
	userservice "apigateway/internal/service/user"
	budgetusecase "apigateway/internal/usecase/budget"
	incomeusecase "apigateway/internal/usecase/income"
	reportusecase "apigateway/internal/usecase/report"
	userusecase "apigateway/internal/usecase/user"
	"log/slog"
)

type App struct {
	HTTPApp *htppapp.App
}

func NewApp(logger *slog.Logger, cfg *config.Config) *App {
	redisDb := redisrepository.NewRedis(*cfg)
	serviceUser := userusecase.NewUserRepo(redisDb, redisDb, redisDb)
	client, err := clientgrpcserver.NewService(cfg)
	if err != nil {
		panic(err)
	}
	userServer := userservice.NewUser(*serviceUser, client, logger, redisDb)
	user := userusecase.NewUserUseCase(userServer)

	budgetService := budgetservice.NewBudget(client, logger)

	budget := budgetusecase.NewBudgetUseCase(budgetService)

	income := incomeservice.NewIncome(logger, client, redisDb)

	incomeUseCase := incomeusecase.NewIncomeUseCase(income)

	report := reportservice.NewReport(logger, client)

	reportUseCase := reportusecase.NewReportUseCase(report)
	server := htppapp.NewApp(logger, cfg.RPCPort, user, *budget, *incomeUseCase, *reportUseCase)
	return &App{
		HTTPApp: server,
	}
}
