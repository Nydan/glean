package http

import (
	"github.com/nydan/glean/internal/controller/http/order"
)

// Controller struct for dependencies
type Controller struct {
	Order *order.Controller
}

// NewController creates http controller
func NewController(o *order.Controller) *Controller {
	return &Controller{
		Order: o,
	}
}
