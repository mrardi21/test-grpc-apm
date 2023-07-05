package main

import (
	"fmt"
	pb "go-proto/pb"
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {

	products := &pb.Products{
		Data: []*pb.Product{
			{
				Id:    1,
				Name:  "Nike Black T-Shirt",
				Price: 10000.00,
				Stock: 100,
				Category: &pb.Category{
					Id:   1,
					Name: "Shirt",
				},
			},
			{
				Id:    1,
				Name:  "Nike Air Jordan",
				Price: 10000.00,
				Stock: 100,
				Category: &pb.Category{
					Id:   2,
					Name: "Shoe",
				},
			},
		},
	}

	data, err := proto.Marshal(products)
	if err != nil {
		log.Fatal("Marshall error")
	}

	// Compact Binary Wire Format
	fmt.Println(data)

	testProducts := &pb.Products{}
	if err = proto.Unmarshal(data, testProducts); err != nil {
		log.Fatal("Unmarshall error", err)
	}

	// fmt.Println(testProducts)

	for _, product := range testProducts.GetData() {
		fmt.Println(product.GetName())
		fmt.Println(product.GetCategory().GetName())
	}

}
