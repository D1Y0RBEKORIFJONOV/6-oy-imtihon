package main

import (
	"budgetservice/budgetproto"
	"budgetservice/budgetservice"
	"budgetservice/internal/storage"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	service := storage.Run()

	budget_service := budgetservice.NewService(*service)
	server := grpc.NewServer()
	budgetproto.RegisterBudgetServiceServer(server, budget_service)
	reflection.Register(server)
	lis, err := net.Listen("tcp", "byudjet-service:8888")
	if err != nil {
		log.Fatal("Unable to listen :", err)
	}
	defer lis.Close()

	log.Println("Server is listening on port ", os.Getenv("server_url"))
	if err = server.Serve(lis); err != nil {
		log.Fatal("Unable to serve :", err)
	}
}
