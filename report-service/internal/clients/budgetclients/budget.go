package budgetclients

import (
	"context"
	"fmt"
	"log"
	"os"
	"report-service/internal/protos/budgetproto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)


func DialBudgetGrpc()budgetproto.BudgetServiceClient{
	conn, err := grpc.NewClient(os.Getenv("budget_url"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to dial grpc client budget:", err)
	}

	return budgetproto.NewBudgetServiceClient(conn)
}

func GetCategories(ctx context.Context, user_id string) (*budgetproto.ListGetUserCategoriesResponse, error){
	listCategories, err := DialBudgetGrpc().GetUserCategories(ctx, &budgetproto.GetUserCategoriesRequest{UserId: user_id})
	if err != nil {
		log.Println("get user categories error: ", err)
		return nil, fmt.Errorf("get user categories error: %v", err)
	}
	return listCategories, nil
}