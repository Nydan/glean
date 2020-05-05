package order

import (
	"fmt"

	"github.com/nydan/glean/internal/entity/order"
	"github.com/pkg/errors"
)

type orderRepoI interface {
	SelectOrderByID(id int) (order.Order, error)
}

// Usecase struct for order domain
type Usecase struct {
	ord orderRepoI
}

// Order adds dependencies to order usecase
func Order(o orderRepoI) *Usecase {
	return &Usecase{
		ord: o,
	}
}

// Create creates new order
func (uc *Usecase) Create() string {
	return "create order from use case layer"
}

// Select order by id
func (uc *Usecase) Select(id int) (order.Order, error) {
	ord := order.Order{}
	if id <= 0 {
		return ord, fmt.Errorf("invalid ID: %d", id)
	}

	ord, err := uc.ord.SelectOrderByID(id)
	if err != nil {
		return ord, errors.Wrap(err, "SelectOrderByID")
	}

	return ord, nil
}
