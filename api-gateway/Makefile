SWAGGER_CMD = swag
GO_RUN_CMD = go run

swagger-init:
	$(SWAGGER_CMD) init -g internal/http/handler/user.go -o internal/app/docs

run:
	$(GO_RUN_CMD) cmd/app/main.go

all: swagger-init run


gen-budget:
	protoc -I protos protos/budget/budget.proto --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative

gen-income:
	protoc -I protos protos/incomeexpences/incomeexpences.proto --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative

gen-report:
	protoc -I protos protos/report/report.proto --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative