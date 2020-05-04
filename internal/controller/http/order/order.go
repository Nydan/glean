package order

import (
	"net/http"
)

type orderUseCase interface {
}

// Controller for order dependency
type Controller struct {
	ouc orderUseCase
}

// Order create new order.Controller
func Order(o orderUseCase) *Controller {
	return &Controller{
		ouc: o,
	}
}

// Create creates order
func (c *Controller) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		w.Write([]byte(`{"data":"create order"}`))
	}
}
