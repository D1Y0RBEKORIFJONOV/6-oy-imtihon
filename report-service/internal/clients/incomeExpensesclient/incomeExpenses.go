package incomeexpensesclient

import (
	"context"
	"fmt"
	"log"
	"os"
	"report-service/internal/protos/incomeExpensesproto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func DialGrpcIncomeExpenses()incomeExpensesproto.IncomeExpensesClient{
	conn, err := grpc.NewClient(os.Getenv("income_url"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to dial grpc client incomeExpenses:", err)
	}

	return incomeExpensesproto.NewIncomeExpensesClient(conn)
}

func GetIncomeExpenses(ctx context.Context, user_id string)(*incomeExpensesproto.GetInfoResponse, error){
	list, err := DialGrpcIncomeExpenses().Info(ctx, &incomeExpensesproto.GetInfoRequest{UserId: user_id})
	if err != nil {
		log.Println("get info error:", err)
		return nil, fmt.Errorf("get info error: %v", err)
	}
	return list, err
}