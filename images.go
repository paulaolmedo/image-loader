package imageLoaderAPI

import (
	"context"
	"errors"
	"fmt"
	images "image-loader/gen/images"
	"image-loader/mongo"
	"io/ioutil"
	"log"
)

type imagesrvc struct {
	logger    *log.Logger
	imagesDao mongo.ImageRepository
}

//NewImageLoader .
func NewImageLoader(logger *log.Logger, imagesDao mongo.ImageRepository) images.Service {
	return &imagesrvc{logger, imagesDao}
}

func (s *imagesrvc) LoadNewSatelliteImage(ctx context.Context, payload *images.RawSatelliteImage) (result *images.GoaResult, err error) {

	filename := *payload.FileName
	if filename == "" {
		return nil, images.MakeBadRequest(errors.New("File name cannot be empty"))
	}

	originalFile, err := ioutil.ReadFile(filename)
	//error reading original file
	if err != nil {
		return nil, images.MakeBadRequest(err)
	}

	bytesWritten, err := s.imagesDao.AddRawImage(originalFile, filename)
	//error adding image to database
	if err != nil {
		return nil, images.MakeErrorAddingImage(err)
	}

	//everything went ok
	code := "200"
	description := fmt.Sprintf("Bytes written %d", bytesWritten)
	result = &images.GoaResult{Code: &code, Description: &description}
	return result, nil
}

func (s *imagesrvc) GetRawSatelliteImage(ctx context.Context, payload *images.GetRawSatelliteImagePayload) (result *images.GoaResult, err error) {
	filename := *payload.FileName
	if filename == "" {
		return nil, images.MakeBadRequest(errors.New("File name cannot be empty"))
	}

	bytesRead, err := s.imagesDao.GetRawImage(filename)

	if err != nil {
		return nil, images.MakeErrorGettingImage(err)
	}

	code := "200"
	description := fmt.Sprintf("Bytes read %d", bytesRead)
	result = &images.GoaResult{Code: &code, Description: &description}
	return result, nil
}

func (s *imagesrvc) LoadNewProcessedSatelliteImage(ctx context.Context, payload *images.ProcessedSatelliteImage) (err error) {
	return nil
}
