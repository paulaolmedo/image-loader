//Copyright (c) 2021 PAULA B. OLMEDO.
//
//This file is part of IMAGE_LOADER
//(see https://github.com/paulaolmedo/image-loader).
//
//This program is free software: you can redistribute it and/or modify
//it under the terms of the GNU General Public License as published by
//the Free Software Foundation, either version 3 of the License, or
//(at your option) any later version.
//
//This program is distributed in the hope that it will be useful,
//but WITHOUT ANY WARRANTY; without even the implied warranty of
//MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//GNU General Public License for more details.
//
//You should have received a copy of the GNU General Public License
//along with this program. If not, see <http://www.gnu.org/licenses/>.

package mongo

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	processed_images "image-loader/internal/models"
	"io/ioutil"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	rawDatabase              = "safo-raw"
	processedImagesDatabase  = "safo-processed-images"
	processedResultsDatabase = "safo-processed-results"
	filesCollection          = "fs.files"
)

// ImageRepository to manage the db access
type ImageRepository interface {
	AddFile(file []byte, filename string, fileType string) (int, error)
	AddProcessedImage(originalFile []byte, Filename string) (int, error) // esta habría que deprecarla
	AddProcessedImageData(image ...*processed_images.ProcessedSatelliteImage) (string, error)
	GetRawImage(Filename string) (int64, error)
	GetProcessedImage(Filename string) (int64, error)
}

// ImageDao to operate with the image objects
type ImageDao struct {
	mongoClient *mongo.Client
}

// InitiateImageDao .
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

// AddRawImage new raw image to the database
func (dao *ImageDao) AddFile(file []byte, filename string, fileType string) (int, error) {
	var err error
	var bucket *gridfs.Bucket

	switch fileType {
	case "raw":
		bucket, err = gridfs.NewBucket(
			dao.mongoClient.Database(rawDatabase),
		)
	case "processed":
		bucket, err = gridfs.NewBucket(
			dao.mongoClient.Database(processedImagesDatabase),
		)
	case "results":
		bucket, err = gridfs.NewBucket(
			dao.mongoClient.Database(processedResultsDatabase),
		)
	default:
		err := errors.New("no such option")
		return 0, err

	}
	// error opening database connection
	if err != nil {
		return 0, err
	}

	uploadStream, err := bucket.OpenUploadStream(
		filename,
	)
	// error opening upload stream
	if err != nil {
		return 0, nil
	}
	defer uploadStream.Close()

	fileSize, err := uploadStream.Write(file)
	// error writing stream
	if err != nil {
		return 0, err
	}
	return fileSize, nil
}

// GetRawImage test
func (dao *ImageDao) GetRawImage(Filename string) (int64, error) {
	db := dao.mongoClient.Database(rawDatabase)
	fsFiles := db.Collection(filesCollection)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var results bson.M
	err := fsFiles.FindOne(ctx, bson.M{}).Decode(&results)
	// error decoding results
	if err != nil {
		return 0, err
	}

	bucket, _ := gridfs.NewBucket(
		db,
	)
	var buf bytes.Buffer
	dStream, err := bucket.DownloadToStreamByName(Filename, &buf)
	// error downloading result
	if err != nil {
		return 0, err
	}

	err = ioutil.WriteFile(Filename, buf.Bytes(), 0o777)
	// error writing file
	if err != nil {
		return 0, err
	}
	// alles gut
	return dStream, nil
}

// AddProcessedImageData .
func (dao *ImageDao) AddProcessedImageData(image ...*processed_images.ProcessedSatelliteImage) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// esto genera una nueva colección en la base de datos de imgs procesadas
	collection := dao.mongoClient.Database(processedResultsDatabase).Collection("data")
	result, err := collection.InsertOne(ctx, bson.M{"data": image})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", result.InsertedID), err
}

// AddProcessedImage new processed image to the database
// TODO parametrizar y que sea un sólo método para las dos cosas
func (dao *ImageDao) AddProcessedImage(originalFile []byte, Filename string) (int, error) {
	bucket, err := gridfs.NewBucket(
		dao.mongoClient.Database(processedImagesDatabase),
	)
	// error opening database connection
	if err != nil {
		return 0, err
	}
	uploadStream, err := bucket.OpenUploadStream(
		Filename,
	)
	// error opening upload stream
	if err != nil {
		return 0, nil
	}
	defer uploadStream.Close()

	fileSize, err := uploadStream.Write(originalFile)
	// error writing stream
	if err != nil {
		return 0, err
	}
	return fileSize, nil
}

// GetProcessedImage test
func (dao *ImageDao) GetProcessedImage(Filename string) (int64, error) {
	db := dao.mongoClient.Database(processedImagesDatabase)
	fsFiles := db.Collection(filesCollection)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var results bson.M
	err := fsFiles.FindOne(ctx, bson.M{}).Decode(&results)
	// error decoding results
	if err != nil {
		return 0, err
	}

	bucket, _ := gridfs.NewBucket(
		db,
	)
	var buf bytes.Buffer
	dStream, err := bucket.DownloadToStreamByName(Filename, &buf)
	// error downloading result
	if err != nil {
		return 0, err
	}

	err = ioutil.WriteFile(Filename, buf.Bytes(), 0o777)
	// error writing file
	if err != nil {
		return 0, err
	}
	// alles gut
	return dStream, nil
}
