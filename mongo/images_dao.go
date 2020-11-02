package mongo

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"

	processed_images "image-loader/gen/processed_images"
)

const (
	rawDatabase       = "safo-raw"
	processedDatabase = "safo-processed"
	filesCollection   = "fs.files"
)

//ImageRepository to manage the db access
type ImageRepository interface {
	AddRawImage(originalFile []byte, filename string) (int, error)
	AddProcessedImage(originalFile []byte, filename string) (int, error)
	AddProcessedImageData(image *processed_images.ProcessedSatelliteImage) (string, error)
	GetRawImage(filename string) (int64, error)
	GetProcessedImage(filename string) (int64, error)
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

//AddProcessedImageData .
func (dao *ImageDao) AddProcessedImageData(imageData *processed_images.ProcessedSatelliteImage) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	collection := dao.mongoClient.Database(processedDatabase).Collection("data")
	result, err := collection.InsertOne(ctx, bson.M{"data": imageData})

	if err != nil { //fmt.Sprintf(bytesWritten)
		return "", err
	}
	return fmt.Sprintf("%s", result.InsertedID), err
}

//AddProcessedImage new processed image to the database
//TODO parametrizar y que sea un sólo método para las dos cosas
func (dao *ImageDao) AddProcessedImage(originalFile []byte, filename string) (int, error) {
	bucket, err := gridfs.NewBucket(
		dao.mongoClient.Database(processedDatabase),
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

//GetProcessedImage test
func (dao *ImageDao) GetProcessedImage(filename string) (int64, error) {

	db := dao.mongoClient.Database(processedDatabase)
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
