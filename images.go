package imageLoaderAPI

import (
	"context"
	"errors"
	"fmt"
	raw_images "image-loader/gen/raw_images"
	"image-loader/mongo"
	"io/ioutil"
	"log"
)

type imagesrvc struct {
	logger    *log.Logger
	imagesDao mongo.ImageRepository
}

//NewImageLoader .
func NewImageLoader(logger *log.Logger, imagesDao mongo.ImageRepository) raw_images.Service {
	return &imagesrvc{logger, imagesDao}
}

func (s *imagesrvc) LoadNewRawSatelliteImage(ctx context.Context, payload *raw_images.RawSatelliteImage) (result *raw_images.GoaResult, err error) {

	filename := *payload.FileName
	if filename == "" {
		return nil, raw_images.MakeBadRequest(errors.New("File name cannot be empty"))
	}

	originalFile, err := ioutil.ReadFile(filename)
	//error reading original file
	if err != nil {
		return nil, raw_images.MakeBadRequest(err)
	}

	bytesWritten, err := s.imagesDao.AddRawImage(originalFile, filename)
	//error adding image to database
	if err != nil {
		return nil, raw_images.MakeErrorAddingImage(err)
	}

	//everything went ok
	code := "200"
	description := fmt.Sprintf("Bytes written %d", bytesWritten)
	result = &raw_images.GoaResult{Code: &code, Description: &description}
	return result, nil
}

func (s *imagesrvc) GetRawSatelliteImage(ctx context.Context, payload *raw_images.GetRawSatelliteImagePayload) (result *raw_images.GoaResult, err error) {
	filename := *payload.FileName
	if filename == "" {
		return nil, raw_images.MakeBadRequest(errors.New("File name cannot be empty"))
	}

	bytesRead, err := s.imagesDao.GetRawImage(filename)

	if err != nil {
		return nil, raw_images.MakeErrorGettingImage(err)
	}

	code := "200"
	description := fmt.Sprintf("Bytes read %d", bytesRead)
	result = &raw_images.GoaResult{Code: &code, Description: &description}
	return result, nil
}
