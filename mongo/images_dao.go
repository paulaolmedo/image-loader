package mongo

import (
	"bytes"
	"context"
	"io/ioutil"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	rawDatabase       = "safo-raw"
	processedDatabase = "safo-processed"
	filesCollection   = "fs.files"
)

//ImageRepository to manage the db access
type ImageRepository interface {
	AddRawImage(originalFile []byte, filename string) (int, error)
	GetRawImage(filename string) (int64, error)
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

//AddRawImage new raw image to the database
func (dao *ImageDao) AddRawImage(originalFile []byte, filename string) (int, error) {
	bucket, err := gridfs.NewBucket(
		dao.mongoClient.Database(rawDatabase),
	)

	//error opening database connection
	if err != nil {
		return 0, err
	}
	uploadStream, err := bucket.OpenUploadStream(
		filename,
	)
	//error opening upload stream
	if err != nil {
		return 0, nil
	}
	defer uploadStream.Close()

	fileSize, err := uploadStream.Write(originalFile)
	//error writing stream
	if err != nil {
		return 0, err
	}
	return fileSize, nil
}

//GetRawImage test
func (dao *ImageDao) GetRawImage(filename string) (int64, error) {

	db := dao.mongoClient.Database(rawDatabase)
	fsFiles := db.Collection(filesCollection)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var results bson.M
	err := fsFiles.FindOne(ctx, bson.M{}).Decode(&results)
	//error decoding results
	if err != nil {
		return 0, err
	}

	bucket, _ := gridfs.NewBucket(
		db,
	)
	var buf bytes.Buffer
	dStream, err := bucket.DownloadToStreamByName(filename, &buf)
	//error downloading result
	if err != nil {
		return 0, err
	}

	err = ioutil.WriteFile(filename, buf.Bytes(), 0777)
	//error writing file
	if err != nil {
		return 0, err
	}
	//alles gut
	return dStream, nil
}
