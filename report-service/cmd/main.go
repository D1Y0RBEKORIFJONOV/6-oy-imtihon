package main

import (
	"log"
	"net"
	"report-service/internal/config"
	"report-service/internal/protos/reportproto"
	reportservice "report-service/reportService"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func main() {
	c := config.Configuration()
	lis, err := net.Listen(c.Report.Host, c.Report.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	reportHandler := &reportservice.ReportGrpcHandler{}
	reportproto.RegisterReportServiceServer(grpcServer, reportHandler)
	reflection.Register(grpcServer)

	log.Println("gRPC server is running on port ", c.Report.Port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
