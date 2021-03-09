package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	run()
}

func run() {
	path := "dataset/order_brush_order.csv"
	file, err := os.Open(path)
	check(err)
	r := csv.NewReader(file)
	count := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		check(err)
		fmt.Println(record)
		count++
		if count == 5 {
			break
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
