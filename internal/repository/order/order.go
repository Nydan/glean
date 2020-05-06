package order

import (
	"github.com/nydan/glean/internal/entity/order"
	"github.com/nydan/glean/pkg/database"
	"github.com/nydan/glean/pkg/redis"
)

// Repository for order
type Repository struct {
	db  database.Database
	rds redis.Redis
}

// Order adds dependencies to order Repository
func Order(db database.Database, rds redis.Redis) *Repository {
	return &Repository{
		db:  db,
		rds: rds,
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
