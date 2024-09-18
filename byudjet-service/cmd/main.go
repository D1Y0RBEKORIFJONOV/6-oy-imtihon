package main

import (
	"budgetservice/budgetproto"
	"budgetservice/budgetservice"
	"budgetservice/internal/config"
	"budgetservice/internal/storage"
	"log"
	"net"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

func main() {
	c := config.Configuration()
	service := storage.Run()

	budget_service := budgetservice.NewService(*service)
	server := grpc.NewServer()
	budgetproto.RegisterBudgetServiceServer(server, budget_service)
	reflection.Register(server)
	lis, err := net.Listen(c.Budget.Host, c.Budget.Port)
	if err != nil {
		log.Fatal("Unable to listen :", err)
	}
	defer lis.Close()

	log.Println("Server is listening on port ", c.Budget.Port)
	if err = server.Serve(lis); err != nil {
		log.Fatal("Unable to serve :", err)
	}
}
