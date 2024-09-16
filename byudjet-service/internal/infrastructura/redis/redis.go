package rediss

// import (
// 	"budgetservice/internal/entity/budget"
// 	"budgetservice/internal/infrastructura/repository"
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"log"

// 	"github.com/google/uuid"
// 	"github.com/redis/go-redis/v9"
// )

// type BudgetRedis struct {
// 	r *redis.Client
// }

// func NewBudgetRedis(r *redis.Client) repository.BudgetRepository {
// 	return &BudgetRedis{r: r}
// }

// func (b *BudgetRedis) AddRedis(ctx context.Context, req budget.CreateBudgetRequest) (*budget.CreateBudgetResponse, error) {
// 	budgetID := uuid.New().String()
// 	count := 0
// 	req.BudgetID = budgetID
// 	req.Spent = 0

// 	budgett, err := b.GetRedis(ctx, req.UserID)
// 	if err != nil {
// 		count++
// 	}

// 	if count == 0 && req.Category == budgett.Category{
// 		log.Println("this category already exsists")
// 		return nil, fmt.Errorf("this category already exsists")
// 	}
// 	jsonValue, err := json.Marshal(req)
// 	if err != nil {
// 		return nil, fmt.Errorf("json formatlashda xatolik: %v", err)
// 	}

// 	err = b.r.Set(ctx, req.UserID, jsonValue, 0).Err()
// 	if err != nil {
// 		return nil, fmt.Errorf("redisga saqlashda xatolik: %v", err)
// 	}

// 	return &budget.CreateBudgetResponse{
// 		Message:  "Byudjet muvaffaqiyatli yaratildi",
// 		BudgetID: budgetID,
// 	}, nil
// }

// func (b *BudgetRedis) GetRedis(ctx context.Context, userID string) (*budget.Budget, error) {
// 	value, err := b.r.Get(ctx, userID).Result()
// 	if err == redis.Nil {
// 		return nil, fmt.Errorf("budget topilmadi")
// 	} else if err != nil {
// 		return nil, fmt.Errorf("redisdan olishda xatolik: %v", err)
// 	}

// 	var budget budget.Budget
// 	err = json.Unmarshal([]byte(value), &budget)
// 	if err != nil {
// 		return nil, fmt.Errorf("json dan obyektga o'tkazishda xatolik: %v", err)
// 	}

// 	return &budget, nil
// }

// func (b *BudgetRedis) GetUserRedis(ctx context.Context, req budget.GetUserCategoriesRequest)(*[]string, error){
// 	values, err := b.r.SMembers(ctx, req.UserID).Result()
// 	if err != nil {
// 		return nil, fmt.Errorf("redis to'plamini olishda xatolik: %v", err)
// 	}
// 	return &values, nil
// }

// func (b *BudgetRedis) UpdateRedis(ctx context.Context, req budget.UpdateBudgetRequest) (*budget.UpdateBudgetResponse, error) {
// 	existingBudget, err := b.GetRedis(ctx, req.UserID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if existingBudget.BudgetID != req.BudgetID {
// 		log.Println("budget id error")
// 		return nil, fmt.Errorf("budget id error")
// 	}

// 	if req.Amount != 0{
// 		existingBudget.Amount = req.Amount
// 	}
// 	if req.Spent != 0{
// 		existingBudget.Spent = req.Spent
// 	}
// 	if req.Currency != ""{
// 		existingBudget.Currency = req.Currency
// 	}
	
// 	jsonValue, err := json.Marshal(existingBudget)
// 	if err != nil {
// 		return nil, fmt.Errorf("json formatlashda xatolik: %v", err)
// 	}

// 	err = b.r.Set(ctx, req.UserID, jsonValue, 0).Err()
// 	if err != nil {
// 		return nil, fmt.Errorf("redisda yangilashda xatolik: %v", err)
// 	}

// 	return &budget.UpdateBudgetResponse{
// 		Message: "Byudjet muvaffaqiyatli yangilandi",
// 	}, nil
// }
