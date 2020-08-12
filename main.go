package main

import (
	"fmt"

	"koko.mongo/mongo"
)

func main() {
	mongo.Init(27017, "koko")
	err := mongo.Connect()
	if err != nil {
		fmt.Println("Failed to connect")
	} else {
		fmt.Println("Connected!")
	}
}
