package router

import (
	"apigateway/internal/http/handler"
	"apigateway/internal/http/middleware"
	budgetusecase "apigateway/internal/usecase/budget"
	incomeusecase "apigateway/internal/usecase/income"
	reportusecase "apigateway/internal/usecase/report"
	userusecase "apigateway/internal/usecase/user"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRouter(user userusecase.User, budget *budgetusecase.BudgetUseCaseImpl,
	income incomeusecase.IncomeUseCaseIml, report reportusecase.ReportUseCaseImpl) *gin.Engine {
	userHandler := handler.NewUserServer(user)
	budgetHandler := handler.NewBudgetHandler(budget)
	incomeHandler := handler.NewIncome(income)
	reportHandler := handler.NewReport(report)
	router := gin.Default()

	router.Use(middleware.Middleware)
	router.Use(middleware.TimingMiddleware)

	url := ginSwagger.URL("swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	userGroup := router.Group("/user")
	{
		userGroup.POST("/login", userHandler.Login)
		userGroup.POST("/register", userHandler.Register)
		userGroup.POST("/verify", userHandler.VerifyUser)
		userGroup.PUT("/update", userHandler.UpdateUser)
		userGroup.PUT("/update/password", userHandler.UpdatePassword)
		userGroup.PUT("/update/email", userHandler.UpdateEmail)
		userGroup.DELETE("/delete", userHandler.DeleteUser)
		userGroup.GET("", userHandler.GetUser)
		userGroup.GET("/all", userHandler.GetAllUser)
	}

	budgetGROUP := router.Group("/budget")
	{
		budgetGROUP.POST("/create", budgetHandler.CreateBudget)
		budgetGROUP.GET("/:category", budgetHandler.GetBudget)
		budgetGROUP.PATCH("/:budget_id", budgetHandler.UpdateBudget)
	}

	incumdeGroup := router.Group("/income")
	{
		incumdeGroup.POST("create/income", incomeHandler.CreateIncome)
		incumdeGroup.GET("/", incomeHandler.GetIncome)
		incumdeGroup.POST("create/expenses", incomeHandler.Expenses)
	}

	reportGroup := router.Group("/report")
	{
		reportGroup.GET("/spending", reportHandler.GetSpending)
		reportGroup.GET("/expenses", reportHandler.GetExpenses)
		reportGroup.GET("/chosen/:start_at/:end_at", reportHandler.GetFromTill)
	}

	return router
}
