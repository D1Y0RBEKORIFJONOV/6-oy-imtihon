package mongodb

import (
	"budgetservice/internal/entity/budget"
	"budgetservice/internal/infrastructura/repository"
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BudgetMongodb struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewBudgetMongodb(client *mongo.Client, collection *mongo.Collection) repository.BudgetRepository {
	return &BudgetMongodb{client: client, collection: collection}
}

func (b *BudgetMongodb) AddMongo(ctx context.Context, req budget.CreateBudgetRequest) (*budget.CreateBudgetResponse, error) {
	budgetID := uuid.New().String()
	req.BudgetID = budgetID
	req.Spent = 0

	_, err := b.collection.InsertOne(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("mongo'ga saqlashda xatolik: %v", err)
	}

	return &budget.CreateBudgetResponse{
		Message:  "Byudjet muvaffaqiyatli yaratildi",
		BudgetID: budgetID,
	}, nil
}

func (b *BudgetMongodb) GetMongo(ctx context.Context, userID, category string) (*budget.Budget, error) {
	filter := bson.M{"userid": userID, "category": category}

	var budget budget.Budget
	err := b.collection.FindOne(ctx, filter).Decode(&budget)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("budget topilmadi")
	} else if err != nil {
		return nil, fmt.Errorf("mongodan olishda xatolik: %v", err)
	}

	return &budget, nil
}

func (b *BudgetMongodb) UpdateMongo(ctx context.Context, req budget.UpdateBudgetRequest) (*budget.UpdateBudgetResponse, error) {
	filter := bson.M{"userid": req.UserID, "budgetid": req.BudgetID}
	log.Println("REQ:", req)
	updateFields := bson.M{}

	if req.Amount != 0 {
		updateFields["amount"] = req.Amount
	}
	if req.Spent != 0 {
		updateFields["spent"] = req.Spent
	}
	if req.Currency != "" {
		updateFields["currency"] = req.Currency
	}

	if len(updateFields) == 0 {
		return nil, fmt.Errorf("hech qanday yangilanish maydoni mavjud emas")
	}

	update := bson.M{"$set": updateFields}

	_, err := b.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("mongo'da yangilashda xatolik: %v", err)
	}

	return &budget.UpdateBudgetResponse{
		Message: "Byudjet muvaffaqiyatli yangilandi",
	}, nil
}

func (b *BudgetMongodb) GetUserMongo(ctx context.Context, req budget.GetUserCategoriesRequest) (*budget.GetBudgetsResponse, error) {
	filter := bson.M{"userid": req.UserID}

	cursor, err := b.collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("foydalanuvchi byudjetlarini olishda xatolik: %v", err)
	}
	defer cursor.Close(ctx)

	var budgets []budget.Budget

	for cursor.Next(ctx) {
		var bgt budget.Budget
		err := cursor.Decode(&bgt)
		if err != nil {
			return nil, fmt.Errorf("byudjetni decode qilishda xatolik: %v", err)
		}
		budgets = append(budgets, bgt)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursorda xatolik: %v", err)
	}

	return &budget.GetBudgetsResponse{
		Budgets: budgets,
	}, nil
}
