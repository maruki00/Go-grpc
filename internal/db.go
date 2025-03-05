package internal

import (
	"fmt"

	orders_order "grpc/protogen/golang/orders"
)

type DB struct {
	collection []*orders_order.Order
}

func NewDB() *DB {
	return &DB{
		collection: make([]*orders_order.Order, 0),
	}
}

func (d *DB) AddOrder(order *orders_order.Order) error {
	for _, o := range d.collection {
		if o.OrderId == order.OrderId {
			return fmt.Errorf("duplicate order id: %d", order.GetOrderId())
		}
	}
	d.collection = append(d.collection, order)
	return nil
}
