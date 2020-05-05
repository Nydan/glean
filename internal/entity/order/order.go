package order

// Order is mapping for table structure from database
type Order struct {
	ID        int     `db:"id"`
	Amount    float64 `db:"amount"`
	ProductID int     `db:"product_id"`
}
