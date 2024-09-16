package main

import (
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"report-service/internal/protos/reportproto"
	reportservice "report-service/reportService"

	"google.golang.org/grpc"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func main() {
	lis, err := net.Listen("tcp", os.Getenv("server_url"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	reportHandler := &reportservice.ReportGrpcHandler{}
	reportproto.RegisterReportServiceServer(grpcServer, reportHandler)
	reflection.Register(grpcServer)

	log.Println("gRPC server is running on port ", os.Getenv("server_url"))
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
