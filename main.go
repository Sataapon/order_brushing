package main

import (
	"fmt"

	"github.com/Sataapon/order_brushing/shop"
)

func main() {
	run()
}

func run() {
	path := "dataset/order_brush_order.csv"
	mapping := shop.New(path)
	_ = mapping
	fmt.Println("complete")
}
