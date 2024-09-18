package storage

import (
	"budgetservice/internal/config"
	"budgetservice/internal/infrastructura/mongodb"
	"budgetservice/internal/service"
	"context"
	"log"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongodb() (*mongo.Client, *mongo.Collection, error) {
	c:=config.Configuration()
	clientOptions := options.Client().ApplyURI(c.Mongo.Url)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, nil, err
	}

	collection := client.Database("Budgets").Collection("budget")

	return client, collection, nil
}

// func NewRedisClient() *redis.Client {
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr:     os.Getenv("redisURL"),
// 		Password: "",
// 		DB:       0,
// 	})
// 	return rdb
// }

func Run() *service.BudgetService {
	// redisclient := NewRedisClient()
	// repo := rediss.NewBudgetRedis(redisclient)
	client, collection, err := NewMongodb()
	if err != nil {
		log.Println("mongodb error")
		return nil
	}
	repo := mongodb.NewBudgetMongodb(client, collection)
	s := service.NewBudgetService(repo)
	return s
}
