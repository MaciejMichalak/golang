package main

import (
	"github.com/MaciejMichalak/golang/mongo/db"
)

func main() {
	db.Init(27017, "koko", nil)
}
