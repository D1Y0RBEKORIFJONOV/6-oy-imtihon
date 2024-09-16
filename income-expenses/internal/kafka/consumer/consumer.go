package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/twmb/franz-go/pkg/kgo"
	"incomeexpenses/internal/protos/income"
	"incomeexpenses/internal/service"
	"log"
)

type Consumer struct {
	Consumer *kgo.Client
	Ctx      context.Context
	Service  service.Service
}

func (u *Consumer) Consume() {
	for {
		fetches := u.Consumer.PollFetches(u.Ctx)
		if err := fetches.Errors(); len(err) > 0 {
			log.Fatal(err)
		}
		fetches.EachPartition(func(ftp kgo.FetchTopicPartition) {
			for _, record := range ftp.Records {
				if err := u.Adjust(record); err != nil {
					log.Println(err)
				}
			}
		})
	}
}

func (u *Consumer) Adjust(record *kgo.Record) error {
	switch string(record.Key) {
	case "incomecreate":
		if err := u.IncomeCreate(record.Value); err != nil {
			log.Println(err)
			return nil
		}
	case "expensescreate":
		if err := u.ExpensesCreate(record.Value); err != nil {
			log.Println(err)
			return err
		}
		return nil
	}
	return nil
}

func (u *Consumer) IncomeCreate(byted []byte) error {
	var req income.CreateIncomeExpensesRequest

	if err := json.Unmarshal(byted, &req); err != nil {
		log.Println(err)
		return err
	}
	_, err := u.Service.Income(u.Ctx, &req)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *Consumer) ExpensesCreate(byted []byte) error {
	var req income.CreateIncomeExpensesRequest
	fmt.Println("req:", string(byted))
	if err := json.Unmarshal(byted, &req); err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("after:", &req)
	_, err := u.Service.Expenses(u.Ctx, &req)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
