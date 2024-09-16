package reportservice

import (
	"context"
	"fmt"
	"log"
	"report-service/internal/clients/budgetclients"
	incomeexpensesclient "report-service/internal/clients/incomeExpensesclient"
	"report-service/internal/protos/reportproto"
	"time"
)

type ReportGrpcHandler struct {
	reportproto.UnimplementedReportServiceServer
}

func (r *ReportGrpcHandler) IncomeExpense(ctx context.Context, req *reportproto.SpendingRequest) (*reportproto.IncomeExpenseResponse, error) {
	var res reportproto.IncomeExpenseResponse
	listCategories, err := budgetclients.GetCategories(ctx, req.UserId)
	if err != nil {
		log.Println("get categories error: ", err)
		return nil, fmt.Errorf("get categories error: %v", err)
	}
	for _, budget := range listCategories.Usercategories {
		res.TotalIncome = res.TotalIncome + budget.Amount
		res.TotalExpenses = res.TotalExpenses + budget.Spent
	}
	res.NetSavings = res.TotalIncome - res.TotalExpenses
	return &res, nil
}

func (r *ReportGrpcHandler) SpendingbyCategory(ctx context.Context, req *reportproto.SpendingRequest) (*reportproto.ListSpendingResponse, error) {
	var spendingres []*reportproto.SpendingResponse
	listCategories, err := budgetclients.GetCategories(ctx, req.UserId)
	if err != nil {
		log.Println("get categories error: ", err)
		return nil, fmt.Errorf("get categories error: %v", err)
	}
	for _, budget := range listCategories.Usercategories {
		var spending reportproto.SpendingResponse
		spending.Category = budget.Category
		spending.Totalspent = budget.Spent
		spendingres = append(spendingres, &spending)
	}

	return &reportproto.ListSpendingResponse{Spent: spendingres}, nil
}

func (r *ReportGrpcHandler) FromTill(ctx context.Context, req *reportproto.FromTillRequest) (*reportproto.FromTillResponse, error) {
	const layout = "02-01-2006"
	var res reportproto.FromTillResponse
	startTime, err := time.Parse(layout, req.StartTime)
	if err != nil {
		log.Println("startTime format xatosi:", err)
		return nil, fmt.Errorf("startTime ni parse qilishda xato: %v", err)
	}
	startTime = startTime.Add(-24 * time.Hour)

	endTime, err := time.Parse(layout, req.EndTime)
	if err != nil {
		log.Println("endTime format xatosi:", err)
		return nil, fmt.Errorf("endTime ni parse qilishda xato: %v", err)
	}
	endTime = endTime.Add(24 * time.Hour)
	listres, err := incomeexpensesclient.GetIncomeExpenses(ctx, req.UserId)
	if err != nil {
		log.Println("get info error:", err)
		return nil, fmt.Errorf("get info error: %v", err)
	}

	for _, v := range listres.Info {
		date, err := time.Parse(layout, v.Date)
		if err != nil {
			log.Println("date format error for item:", v.Date)
			continue
		}
		fmt.Println("data:", date, "start:", startTime, "end:", endTime)

		if date.After(startTime) && date.Before(endTime) {
			fmt.Println(">>>>>>", v.Amount)
			res.TotalAmount = res.TotalAmount + v.Amount
		}
	}

	return &reportproto.FromTillResponse{TotalAmount: res.TotalAmount}, nil
}
