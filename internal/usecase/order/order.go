package order

type orderRepoI interface {
}

// Usecase struct for order domain
type Usecase struct {
	ord orderRepoI
}

// NewOrder creates new order Usecase
func NewOrder(o orderRepoI) *Usecase {
	return &Usecase{
		ord: o,
	}
}

// Create creates new order
func (uc *Usecase) Create() string {
	return "create order from use case layer"
}
