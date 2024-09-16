package connections

import (
	"context"
	"database/sql"
	"fmt"
	budgetservice "incomeexpenses/internal/clients/budget"
	"incomeexpenses/internal/clients/notification"
	"incomeexpenses/internal/config"
	storage "incomeexpenses/internal/database"
	interface17 "incomeexpenses/internal/interface"
	"incomeexpenses/internal/kafka/consumer"
	"incomeexpenses/internal/service"
	"incomeexpenses/internal/service/adjust"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/twmb/franz-go/pkg/kgo"
)

func Database() *sql.DB {
	c := config.Configuration()

	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", c.Database.User, c.Database.Password, c.Database.Host, c.Database.Port, c.Database.DBname))
	if err != nil {
		log.Println(err)
		return nil
	}
	if err := db.Ping(); err != nil {
		log.Println(err)
		return nil
	}
	return db
}

func Adjust() interface17.IncomeExpenses {
	ctx := context.Background()
	db := Database()
	notification := notification.Notification()
	queries := storage.New(db)
	budget := budgetservice.Budget()
	return &adjust.Adjust{Ctx: ctx, Sql: *queries, Budget: budget, Notification: notification}
}

func Service() *service.Service {
	a := Adjust()
	return &service.Service{S: a}
}

func Consumer() *consumer.Consumer {
	c := config.Configuration()
	var (
		err      error
		consmr *kgo.Client
	)
	for i := 0; i < 1; i++ {
		consmr, err = kgo.NewClient(
			kgo.SeedBrokers(c.Kafka.Port),
			kgo.ConsumeTopics(c.Kafka.Topic),
			kgo.ConsumerGroup("income_service"),
		)
		if err != nil {
			log.Println(err)
			time.Sleep(1 * time.Millisecond)
			continue
		}
		break
	}
	ctx := context.Background()

	service1 := Service()

	return &consumer.Consumer{Consumer: consmr, Ctx: ctx, Service: *service1}
}
