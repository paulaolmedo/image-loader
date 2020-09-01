package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	rawDatabase       = "safo-raw"
	processedDatabase = "safo-processed"
	filesCollection   = "fs.files"
)

//ImageRepository to manage the db access
type ImageRepository interface {
}

//ImageDao to operate with the image objects
type ImageDao struct {
	mongoClient *mongo.Client
}

//InitiateImageDao .
func InitiateImageDao(url string) (*ImageDao, error) {
	opts := options.Client()
	opts.ApplyURI(url)
	opts.SetMaxPoolSize(5)

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}
	dao := ImageDao{mongoClient: client}
	return &dao, nil
}
