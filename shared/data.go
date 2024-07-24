package shared

// Expense represents data about an expense.
type Expense struct {
	ID          int64   `json:"id"`
	Amount      float64 `json:"amount"`
	CategoryID  int64   `json:"category_id"`
	Description string  `json:"description"`
	Epoch       float64 `json:"epoch"`
}
