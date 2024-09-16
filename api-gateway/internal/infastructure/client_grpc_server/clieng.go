package clientgrpcserver

import (
	budgetproto "apigateway/gen/go/budget"
	incomepb "apigateway/gen/go/incomeexpences"
	reportproto "apigateway/gen/go/report"
	"apigateway/internal/config"
	"fmt"
	user1 "github.com/D1Y0RBEKORIFJONOV/SmartHome_Protos/gen/go/user"
	notificationpb "github.com/D1Y0RBEKORIFJONOV/ekzamen-5protos/gen/go/notification"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type ServiceClient interface {
	UserServiceClient() user1.UserServiceClient
	NotificationServiceClient() notificationpb.NotificationServiceClient
	BudgetServiceClient() budgetproto.BudgetServiceClient
	IncomeServiceClient() incomepb.IncomeExpensesClient
	ReportServiceClient() reportproto.ReportServiceClient
	Close() error
}

type serviceClient struct {
	connection                []*grpc.ClientConn
	userService               user1.UserServiceClient
	budgetService             budgetproto.BudgetServiceClient
	incomeExpenseService      incomepb.IncomeExpensesClient
	notificationServiceClient notificationpb.NotificationServiceClient
	reportClient              reportproto.ReportServiceClient
}

func NewService(cfg *config.Config) (ServiceClient, error) {
	connSoldiersService, err := grpc.NewClient(fmt.Sprintf("%s", cfg.UserUrl),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	connnotificationServiceClient, err := grpc.NewClient(fmt.Sprintf("%s", cfg.NotificationUrl),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	conBudgetServiceClient, err := grpc.NewClient(fmt.Sprintf("%s", cfg.BudgetServiceUrl),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	incomeExpenseServiceConn, err := grpc.NewClient(fmt.Sprintf("%s", cfg.IncomeServiceUrl),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	reportConn, err := grpc.NewClient(fmt.Sprintf("%s", cfg.ReportUcl),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &serviceClient{
		userService:               user1.NewUserServiceClient(connSoldiersService),
		notificationServiceClient: notificationpb.NewNotificationServiceClient(connnotificationServiceClient),
		budgetService:             budgetproto.NewBudgetServiceClient(conBudgetServiceClient),
		incomeExpenseService:      incomepb.NewIncomeExpensesClient(incomeExpenseServiceConn),
		reportClient:              reportproto.NewReportServiceClient(reportConn),
		connection:                []*grpc.ClientConn{connSoldiersService, connnotificationServiceClient, conBudgetServiceClient, incomeExpenseServiceConn, reportConn},
	}, nil
}

func (s *serviceClient) UserServiceClient() user1.UserServiceClient {
	return s.userService
}
func (s *serviceClient) NotificationServiceClient() notificationpb.NotificationServiceClient {
	return s.notificationServiceClient
}
func (s *serviceClient) IncomeServiceClient() incomepb.IncomeExpensesClient {
	return s.incomeExpenseService
}

func (s *serviceClient) BudgetServiceClient() budgetproto.BudgetServiceClient {
	return s.budgetService
}

func (s *serviceClient) ReportServiceClient() reportproto.ReportServiceClient {
	return s.reportClient
}

func (s *serviceClient) Close() error {
	var err error
	for _, conn := range s.connection {
		if cer := conn.Close(); cer != nil {
			log.Println("Error while closing gRPC connection:", cer)
			err = cer
		}
	}
	return err
}
