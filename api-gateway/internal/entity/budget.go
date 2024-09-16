package entity

type (
	CreateBudgetRequest struct {
		UserID   string  `json:"-"`
		Category string  `json:"category"`
		Amount   float32 `json:"amount"`
		Currency string  `json:"currency"`
	}

	CreateBudgetResponse struct {
		Message  string `json:"message"`
		BudgetID string `json:"budget_id"`
	}

	GetBudgetsRequest struct {
		UserID   string `json:"-"`
		Category string `json:"category"`
	}

	GetBudgetsResponse struct {
		Budgets Budget `json:"budgets"`
	}

	Budget struct {
		BudgetID string  `json:"budget_id"`
		Category string  `json:"category"`
		Amount   float32 `json:"amount"`
		Spent    float32 `json:"spent"`
		Currency string  `json:"currency"`
		UserID   string  `json:"-"`
	}

	UpdateBudgetRequest struct {
		BudgetID string  `json:"-"`
		Amount   float32 `json:"amount"`
		UserID   string  `json:"-"`
		Spent    float32 `json:"spent"`
		Currency string  `json:"currency"`
	}

	UpdateBudgetResponse struct {
		Message string `json:"message"`
	}

	GetUserCategoriesRequest struct {
		UserID string `json:"-"`
	}

	ListGetUserCategoriesResponse struct {
		UserCategories []Budget `json:"usercategories"`
	}
)
