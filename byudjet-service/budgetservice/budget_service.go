package budgetservice

import (
	"budgetservice/budgetproto"
	"budgetservice/internal/entity/budget"
	"budgetservice/internal/service"
	"context"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	budgetproto.UnimplementedBudgetServiceServer
	s service.BudgetService
}

func NewService(s service.BudgetService) *Service {
	return &Service{s: s}
}

func (s *Service) CreateBudget(ctx context.Context, req *budgetproto.CreateBudgetRequest) (*budgetproto.CreateBudgetResponse, error) {

	_, err := s.GetBudgets(ctx, &budgetproto.GetBudgetsRequest{
		UserId:   req.UserId,
		Category: req.Category,
	})

	if err != nil && errors.Is(err, mongo.ErrNoDocuments) {
		return nil, fmt.Errorf("xatolik yuz berdi: %v", err)
	}

	if err == nil {
		return nil, fmt.Errorf("category allaqachon mavjud: %v", err)
	}

	res, err := s.s.Createbyudjet(ctx, budget.CreateBudgetRequest{
		UserID:   req.UserId,
		Category: req.Category,
		Currency: req.Currency,
		Amount:   req.Amount,
	})
	if err != nil {
		log.Println("budjet yaratishda xatolik:", err)
		return nil, fmt.Errorf("budjet yaratishda xatolik: %v", err)
	}

	return &budgetproto.CreateBudgetResponse{Message: res.Message, BudgetId: res.BudgetID}, nil
}



func (s *Service) GetBudgets(ctx context.Context, req *budgetproto.GetBudgetsRequest) (*budgetproto.GetBudgetsResponse, error) {
	var budgetres budgetproto.Budget
	res, err := s.s.Getbyudjet(ctx, req.UserId, req.Category)
	if err != nil {
		log.Println("getbudget error:", err)
		return nil, fmt.Errorf("getbudget error:, %v", err)
	}

	budgetres.UserId = res.UserID
	budgetres.BudgetId = res.BudgetID
	budgetres.Amount = res.Amount
	budgetres.Spent = res.Spent
	budgetres.Category = res.Category
	budgetres.Currency = res.Currency
	return &budgetproto.GetBudgetsResponse{Budgets: &budgetres}, nil
}

func (s *Service) UpdateBudget(ctx context.Context,req *budgetproto.UpdateBudgetRequest) (*budgetproto.UpdateBudgetResponse, error){
	res, err := s.s.Updatebudget(ctx, budget.UpdateBudgetRequest{
		BudgetID: req.BudgetId,
		Currency: req.Currency,
		UserID: req.UserId,
		Amount: req.Amount,
		Spent: req.Spent,
	})
	if err != nil {
		log.Println("update budget error")
		return nil, fmt.Errorf("update budget error: %v", err)
	}

	return &budgetproto.UpdateBudgetResponse{Message: res.Message}, nil
}

func (s *Service) GetUserCategories(ctx context.Context, req *budgetproto.GetUserCategoriesRequest) (*budgetproto.ListGetUserCategoriesResponse, error) {
    res, err := s.s.Getusers(ctx, budget.GetUserCategoriesRequest{UserID: req.UserId})
    if err != nil {
        log.Println("getusers error:", err)
        return nil, fmt.Errorf("getusers error: %v", err)
    }

    // Kategoriyalarni qaytarish
    var categories []*budgetproto.Budget
    for _, budget := range res.Budgets {
        categories = append(categories, &budgetproto.Budget{
				BudgetId: budget.BudgetID,
				Category: budget.Category,
				Amount: budget.Amount,
				Spent: budget.Spent,
				Currency: budget.Currency,
				UserId: budget.UserID,
        })
    }

    return &budgetproto.ListGetUserCategoriesResponse{Usercategories: categories}, nil
}
