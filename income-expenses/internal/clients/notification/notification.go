package notification

import (
	"fmt"
	"incomeexpenses/internal/config"
	"log"

	notificationpb "github.com/D1Y0RBEKORIFJONOV/ekzamen-5protos/gen/go/notification"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Notification() notificationpb.NotificationServiceClient {
	c := config.Configuration()
	conn, err := grpc.NewClient(fmt.Sprintf("%v", c.Notification.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
	}
	client := notificationpb.NewNotificationServiceClient(conn)
	return client
}
