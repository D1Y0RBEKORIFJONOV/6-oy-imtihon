package htppapp

import (
	"apigateway/internal/http/router"
	budgetusecase "apigateway/internal/usecase/budget"
	incomeusecase "apigateway/internal/usecase/income"
	reportusecase "apigateway/internal/usecase/report"
	userusecase "apigateway/internal/usecase/user"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type App struct {
	Logger *slog.Logger
	Port   string
	Server *gin.Engine
}

func NewApp(logger *slog.Logger, port string, handlerService userusecase.User,
	budget budgetusecase.BudgetUseCaseImpl,
	income incomeusecase.IncomeUseCaseIml, report reportusecase.ReportUseCaseImpl) *App {
	sever := router.RegisterRouter(handlerService, &budget, income, report)
	return &App{
		Port:   port,
		Server: sever,
		Logger: logger,
	}
}

func (app *App) Start() {
	const op = "app.Start"
	log := app.Logger.With(
		slog.String(op, "Starting server"),
		slog.String("port", app.Port))
	log.Info("Starting server")
	err := app.Server.SetTrustedProxies(nil)
	if err != nil {
		log.Error("Error setting trusted proxies", "error", err)
		return
	}
	err = app.Server.RunTLS(app.Port,
		"/localhost.pem",
		"/localhost-key.pem")
	if err != nil {
		log.Error("Error starting server", "error", err)
		return
	}
}
