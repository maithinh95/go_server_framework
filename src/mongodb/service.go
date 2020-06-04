package mongodb

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type service struct {
	database *mongo.Database
}

/*
Service repository
*/
type Service interface {
}

/*
NewService create a new repository
*/
func NewService(uri, dbname string) Service {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		logrus.Fatal(err)
	}
	return &service{
		database: client.Database(dbname),
	}
}
