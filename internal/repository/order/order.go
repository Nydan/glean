package order

import (
	"github.com/jmoiron/sqlx"
	"github.com/nydan/glean/internal/entity/order"
)

// Repository for order
type Repository struct {
	db *sqlx.DB
}

// Order adds dependencies to order Repository
func Order(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// SelectOrderByID selects order data by id.
// Return empty Order struct when no record found.
func (r *Repository) SelectOrderByID(id int) (order.Order, error) {
	ord := order.Order{}

	err := r.db.Select(&ord, "SELECT * FROM order WHERE id = $1", id)
	if err != nil {
		return ord, err
	}

	return ord, nil
}
