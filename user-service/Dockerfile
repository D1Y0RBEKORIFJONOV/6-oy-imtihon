FROM golang:1.23.0-alpine AS build

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /src/cmd/user

RUN go build -o /bin/user

FROM alpine:latest

RUN apk add --no-cache ca-certificates

COPY --from=build /bin/user /bin/user

CMD ["/bin/user"]
