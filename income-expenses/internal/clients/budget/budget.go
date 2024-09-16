package budgetservice

import (
	"fmt"
	"incomeexpenses/internal/config"
	"incomeexpenses/internal/protos/budgetproto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Budget() budgetproto.BudgetServiceClient {
	c := config.Configuration()
	conn, err := grpc.NewClient(fmt.Sprintf("%v", c.Budget.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
	}
	client := budgetproto.NewBudgetServiceClient(conn)
	return client
}
