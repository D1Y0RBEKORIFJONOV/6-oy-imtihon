package producer

import (
	"apigateway/internal/config"
	"context"
	"log"

	"github.com/twmb/franz-go/pkg/kgo"
)

func Producer(ctx context.Context, key []byte, req []byte, topicName string) error {
	cfg := config.New()
	client, err := kgo.NewClient(
		kgo.SeedBrokers(cfg.MessageBrokerUses.URL),
		kgo.AllowAutoTopicCreation(),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	if err := client.Ping(ctx); err != nil {
		log.Println("client not connected to kafka", err)
	}

	record := kgo.Record{
		Key:   key,
		Topic: topicName,
		Value: req,
	}
	if err := client.ProduceSync(ctx, &record).FirstErr(); err != nil {
		log.Println(err)
	}
	return nil
}
