package order

import (
	"encoding/json"
	"net/http"

	"github.com/nydan/glean/internal/entity/order"
	"github.com/nydan/glean/pkg/logger"
)

type orderUseCase interface {
	Create() string
	Select(id int) (order.Order, error)
}

// Controller for order dependency
type Controller struct {
	ouc orderUseCase
}

// Order adds dependencies to order controller
func Order(o orderUseCase) *Controller {
	return &Controller{
		ouc: o,
	}
}

// Create creates order
func (c *Controller) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)

		data := c.ouc.Create()

		resp := map[string]interface{}{
			"data": data,
		}

		buf, err := json.Marshal(resp)
		if err != nil {
			logger.Errorw("failed marshal", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.Write(buf)
	}
}
