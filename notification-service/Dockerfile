FROM golang:1.23.0-alpine AS build

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /src/cmd/notification

RUN go build -o /bin/ntfk

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=build /bin/ntfk /bin/ntfk

CMD ["/bin/ntfk"]
