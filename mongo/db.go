package mongo

import (
	"context"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoConfig struct {
	port   int
	dbname string
	client *mongo.Client
}

var config mongoConfig = mongoConfig{27017, "koko", nil}

func Init(port int, dbname string) {
	config.port = port
	config.dbname = dbname
}

func Connect() error {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:" + strconv.Itoa(config.port))
	var err error = nil
	config.client, err = mongo.Connect(context.TODO(), clientOptions)
	return err
}
