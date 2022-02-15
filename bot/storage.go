package bot

type Storage interface {
	GetOrders() ([]Order, error)
	SaveOrder(Order) error
	DeleteOrder(Order) error
}
