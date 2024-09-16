package budget

type CreateBudgetRequest struct {
    BudgetID string `json:"budget_id"`
    UserID  string  `json:"user_id"`
    Category string  `json:"category"`
    Amount   float32 `json:"amount"`
    Currency string  `json:"currency"`
    Spent    float32 `json:"spent"`
}

type CreateBudgetResponse struct {
    Message  string `json:"message"`
    BudgetID string `json:"budget_id"`
}

type GetBudgetsRequest struct {
    UserID string `json:"user_id"`
    Category string  `json:"category"`
}

type GetBudgetsResponse struct {
    Budgets []Budget `json:"budgets"`
}

type Budget struct {
    BudgetID string  `json:"budget_id"`
    Category string  `json:"category"`
    Amount   float32 `json:"amount"`
    Spent    float32 `json:"spent"`
    Currency string  `json:"currency"`
    UserID  string  `json:"user_id"`
}

type UpdateBudgetRequest struct {
    BudgetID string  `json:"budget_id"`
    UserID  string  `json:"user_id"`
    Amount   float32 `json:"amount"`
    Spent    float32 `json:"spent"`
    Currency string  `json:"currency"`
}

type UpdateBudgetResponse struct {
    Message string `json:"message"`
}

type GetUserCategoriesRequest struct{
    UserID  string  `json:"user_id"`
}