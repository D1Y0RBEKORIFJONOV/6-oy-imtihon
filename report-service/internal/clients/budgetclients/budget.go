package budgetclients

import (
	"context"
	"fmt"
	"log"
	"report-service/internal/config"
	"report-service/internal/protos/budgetproto"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialBudgetGrpc() budgetproto.BudgetServiceClient {
	c := config.Configuration()
	conn, err := grpc.NewClient(c.Budget.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to dial grpc client budget:", err)
	}

	return budgetproto.NewBudgetServiceClient(conn)
}

func GetCategories(ctx context.Context, user_id string) (*budgetproto.ListGetUserCategoriesResponse, error) {
	listCategories, err := DialBudgetGrpc().GetUserCategories(ctx, &budgetproto.GetUserCategoriesRequest{UserId: user_id})
	if err != nil {
		log.Println("get user categories error: ", err)
		return nil, fmt.Errorf("get user categories error: %v", err)
	}
	return listCategories, nil
}
