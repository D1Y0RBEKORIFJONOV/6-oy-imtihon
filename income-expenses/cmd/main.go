package main

import (
	"fmt"
	"incomeexpenses/internal/config"
	"incomeexpenses/internal/connections"
	"incomeexpenses/internal/protos/income"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	c := config.Configuration()
	ls, err := net.Listen(c.IncomeExpenses.Host, c.IncomeExpenses.Port)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	service := connections.Service()
	income.RegisterIncomeExpensesServer(s, service)
	reflection.Register(s)
	fmt.Printf("server started on the port %s\n", c.IncomeExpenses.Port)

	go func() {
		consumer:=connections.Consumer()

		consumer.Consume()
	}()

	if err := s.Serve(ls); err != nil {
		log.Fatal(err)
	}
}
