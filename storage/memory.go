package storage

import "github.com/fedealliani/gobot/bot"

type memoryStorage struct {
	m map[int64]bot.Order
}

func CreateMemoryStorage() *memoryStorage {
	m := map[int64]bot.Order{}
	return &memoryStorage{m: m}
}

func (m *memoryStorage) GetOrders() ([]bot.Order, error) {
	orders := []bot.Order{}
	for _, v := range m.m {
		orders = append(orders, v)
	}
	return orders, nil
}

func (m *memoryStorage) SaveOrder(order bot.Order) error {
	m.m[order.ID] = order
	return nil
}

func (m *memoryStorage) DeleteOrder(id int64) error {
	delete(m.m, id)
	return nil
}
