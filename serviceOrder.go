package main

import (
	"context"

	orders_order "grpc/protogen/golang/orders"
	"google.golang.org/grpc"
)

type OrdersClient interface {
	AddOrder(ctx context.Context, in *orders_order.PayloadWithSingleOrder, opts ...grpc.CallOption) (*orders_order.Empty, error)
}
type OrderSErvice struct {
}
