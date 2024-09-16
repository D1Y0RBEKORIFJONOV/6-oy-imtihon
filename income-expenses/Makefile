DB="postgres://postgres:+_+diyor2005+_+@localhost:5432/postgres?sslmode=disable"

run:
	go run cmd/main.go

table:
	migrate create -dir migrations -ext sql db
table_up:
	migrate -path migrations -database ${DB} up
table_down:
	migrate -path migrations -database ${DB} down
table_force:
	migrate -path migrations -database ${DB} force 1
gen-proto:
	@protoc \
	--go_out=. \
	--go-grpc_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	./internal/protos/IncomeExpenses.proto