package adjust

import (
	"context"
	"database/sql"
	"fmt"
	storage "incomeexpenses/internal/database"
	"incomeexpenses/internal/protos/budgetproto"
	"incomeexpenses/internal/protos/income"
	"log"

	notificationpb "github.com/D1Y0RBEKORIFJONOV/ekzamen-5protos/gen/go/notification"
	"github.com/google/uuid"
)

type Adjust struct {
	Ctx          context.Context
	Sql          storage.Queries
	Budget       budgetproto.BudgetServiceClient
	Notification notificationpb.NotificationServiceClient
}

func (u *Adjust) Income(ctx context.Context, req *income.CreateIncomeExpensesRequest) (*income.CreateIncomeExpensesResponse, error) {
	var newreq = storage.InsertInfoParams{
		ID:       sql.NullString{String: uuid.New().String(), Valid: true},
		Userid:   sql.NullString{String: req.UserId, Valid: true},
		Type:     sql.NullString{String: req.Type, Valid: true},
		Category: sql.NullString{String: req.Category, Valid: true},
		Currency: sql.NullString{String: req.Currency, Valid: true},
		Amount:   sql.NullFloat64{Float64: float64(req.Amount), Valid: true},
		Date:     sql.NullString{String: req.Date, Valid: true},
	}
	res, err := u.Sql.InsertInfo(ctx, newreq)
	if err != nil {
		_, err = u.Notification.AddNotification(u.Ctx, &notificationpb.AddNotificationReq{
			UserId: req.UserId,
			Messages: &notificationpb.CreateMessage{
				SenderName: "income-expenses",
				Status:     err.Error()}})
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return nil, err
	}
	_, err = u.Notification.AddNotification(u.Ctx, &notificationpb.AddNotificationReq{
		UserId: req.UserId,
		Messages: &notificationpb.CreateMessage{
			SenderName: "income-expenses",
			Status:     fmt.Sprintf("Note has been saved to %v category with this %v sum and it's id is %v\n", req.Category, req.Amount, res.String)}})
	if err != nil {
		_, err = u.Notification.AddNotification(u.Ctx, &notificationpb.AddNotificationReq{
			UserId: req.UserId,
			Messages: &notificationpb.CreateMessage{
				SenderName: "income-expenses",
				Status:     err.Error()}})
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return nil, err
	}
	return &income.CreateIncomeExpensesResponse{Message: "Note hase been saved", Transactionid: res.String}, nil
}

func (u *Adjust) Expenses(ctx context.Context, req *income.CreateIncomeExpensesRequest) (*income.CreateIncomeExpensesResponse, error) {
	log.Println(req.UserId)
	budget, err := u.Budget.GetBudgets(u.Ctx, &budgetproto.GetBudgetsRequest{UserId: req.UserId, Category: req.Category})
	if err != nil {
		_, err = u.Notification.AddNotification(u.Ctx, &notificationpb.AddNotificationReq{
			UserId: req.UserId,
			Messages: &notificationpb.CreateMessage{
				SenderName: "income-expenses",
				Status:     err.Error()}})
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	res, err := u.MoneyChecker(budget, req)
	if err != nil {
		_, err = u.Notification.AddNotification(u.Ctx, &notificationpb.AddNotificationReq{
			UserId: req.UserId,
			Messages: &notificationpb.CreateMessage{
				SenderName: "income-expenses",
				Status:     err.Error()}})
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return nil, err
	}
	return &income.CreateIncomeExpensesResponse{Message: "note has been saved", Transactionid: res}, nil
}

func (u *Adjust) Get(ctx context.Context, req *income.GetInfoRequest) (*income.GetInfoResponse, error) {
	res, err := u.Sql.Getinfo(ctx, sql.NullString{String: req.UserId, Valid: true})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var res1 []*income.InfoResponse

	for _, v := range res {
		var all = income.InfoResponse{
			Transactionid: v.ID.String,
			Type:          v.Type.String,
			UserId:        v.Userid.String,
			Category:      v.Category.String,
			Currency:      v.Currency.String,
			Amount:        float32(v.Amount.Float64),
			Date:          v.Date.String,
		}
		res1 = append(res1, &all)
	}
	return &income.GetInfoResponse{Info: res1}, nil
}

func (u *Adjust) MoneyChecker(req *budgetproto.GetBudgetsResponse, req1 *income.CreateIncomeExpensesRequest) (string, error) {
	switch {
	case req1.Amount > req.Budgets.Amount:
		_, err := u.Notification.AddNotification(u.Ctx, &notificationpb.AddNotificationReq{
			UserId: req1.UserId,
			Messages: &notificationpb.CreateMessage{
				SenderName: "income-expenses",
				Status:     "The expenses are getting more that a budget that supposed to use for this category"}})
		if err != nil {
			_, err = u.Notification.AddNotification(u.Ctx, &notificationpb.AddNotificationReq{
				UserId: req1.UserId,
				Messages: &notificationpb.CreateMessage{
					SenderName: "income-expenses",
					Status:     err.Error()}})
			if err != nil {
				log.Println(err)
				return "", nil
			}
			return "", nil
		}
		return "", nil
	case req1.Amount == req.Budgets.Amount:
		var newreq = storage.InsertInfoParams{
			ID:       sql.NullString{String: uuid.New().String(), Valid: true},
			Userid:   sql.NullString{String: req1.UserId, Valid: true},
			Type:     sql.NullString{String: req1.Type, Valid: true},
			Category: sql.NullString{String: req1.Category, Valid: true},
			Currency: sql.NullString{String: req1.Currency, Valid: true},
			Amount:   sql.NullFloat64{Float64: float64(req1.Amount), Valid: true},
			Date:     sql.NullString{String: req1.Date, Valid: true},
		}
		res, err := u.Sql.InsertInfo(u.Ctx, newreq)
		if err != nil {
			_, err = u.Notification.AddNotification(u.Ctx, &notificationpb.AddNotificationReq{
				UserId: req1.UserId,
				Messages: &notificationpb.CreateMessage{
					SenderName: "income-expenses",
					Status:     err.Error()}})
			if err != nil {
				log.Println(err)
				return "", nil
			}
			return "", err
		}
		req.Budgets.Amount -= req.Budgets.Spent
		_, err = u.Budget.UpdateBudget(u.Ctx, &budgetproto.UpdateBudgetRequest{BudgetId: req.Budgets.BudgetId, UserId: req1.UserId, Spent: req1.Amount + req.Budgets.Spent, Amount: req.Budgets.Amount})
		if err != nil {
			log.Println(err)
			return "", err
		}
		_, err = u.Notification.AddNotification(u.Ctx, &notificationpb.AddNotificationReq{
			UserId: req1.UserId,
			Messages: &notificationpb.CreateMessage{
				SenderName: "income-expenses",
				Status:     "You have no left any pounds for this type of the expenses"}})
		if err != nil {
			log.Println(err)
			return "", nil
		}
		return res.String, nil
	case req1.Amount < req.Budgets.Amount:
		var newreq = storage.InsertInfoParams{
			ID:       sql.NullString{String: uuid.New().String(), Valid: true},
			Userid:   sql.NullString{String: req1.UserId, Valid: true},
			Type:     sql.NullString{String: req1.Type, Valid: true},
			Category: sql.NullString{String: req1.Category, Valid: true},
			Currency: sql.NullString{String: req1.Currency, Valid: true},
			Amount:   sql.NullFloat64{Float64: float64(req1.Amount), Valid: true},
			Date:     sql.NullString{String: req1.Date, Valid: true},
		}
		res, err := u.Sql.InsertInfo(u.Ctx, newreq)
		if err != nil {
			log.Println(err)
			return "", err
		}
		req.Budgets.Amount -= req1.Amount
		if req.Budgets.Amount == 0 {
			_, err = u.Notification.AddNotification(u.Ctx, &notificationpb.AddNotificationReq{
				UserId: req1.UserId,
				Messages: &notificationpb.CreateMessage{
					SenderName: "income-expenses",
					Status:     "you don't have money left for this type of category"}})
			if err != nil {
				log.Println(err)
				return "", nil
			}
		}
		_, err = u.Budget.UpdateBudget(u.Ctx, &budgetproto.UpdateBudgetRequest{BudgetId: req.Budgets.BudgetId, UserId: req1.UserId, Spent: req.Budgets.Spent + req1.Amount, Amount: req.Budgets.Amount})
		if err != nil {
			log.Println(err)
			return "", nil
		}
		_, err = u.Notification.AddNotification(u.Ctx, &notificationpb.AddNotificationReq{
			UserId: req1.UserId,
			Messages: &notificationpb.CreateMessage{
				SenderName: "income-expenses",
				Status:     fmt.Sprintf("Note has been saved to %v category with this %v sum and its id is %v. You have left %v money", req1.Category, req1.Amount, res.String, req.Budgets.Amount)}})
		if err != nil {
			log.Println(err)
			return "", err
		}
		return res.String, nil

	}
	return "", nil

}
