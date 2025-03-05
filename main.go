// main.go

package main

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/encoding/protojson"

	orders_order "grpc/protogen/golang/orders"
)

func main() {
	orderItem := orders_order.Order{
		OrderId:    10,
		CustomerId: 11,
		IsActive:   true,
		OrderDate:  "01/01/2021",
		Products: []*orders_order.Product{
			{ProductId: 1, ProductName: "CocaCola", ProductType: orders_order.ProductType_DRINK},
		},
	}

	bytes, err := protojson.Marshal(&orderItem)
	if err != nil {
		log.Fatal("deserialization error:", err)
	}

	fmt.Println(string(bytes))
}
