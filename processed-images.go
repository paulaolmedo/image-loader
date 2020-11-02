package imageloaderapi

import (
	"context"
	"errors"
	"fmt"
	processed_images "image-loader/gen/processed_images"
	"image-loader/mongo"
	"io/ioutil"
	"log"
)

type processedimagessrvc struct {
	logger    *log.Logger
	imagesDao mongo.ImageRepository
}

//NewProcessedImageLoader .
func NewProcessedImageLoader(logger *log.Logger, imagesDao mongo.ImageRepository) processed_images.Service {
	return &processedimagessrvc{logger, imagesDao}
}

func (s *processedimagessrvc) LoadNewProcessedSatelliteImage(ctx context.Context, payload *processed_images.ProcessedSatelliteImage) (*processed_images.GoaResult, error) {
	filename := *payload.FileName
	if filename == "" {
		return nil, processed_images.MakeBadRequest(errors.New("File name cannot be empty"))
	}

	originalFile, err := ioutil.ReadFile(filename)
	//error reading original file
	if err != nil {
		return nil, processed_images.MakeBadRequest(err)
	}

	bytesWritten, err := s.imagesDao.AddProcessedImage(originalFile, filename)
	//error adding image to database
	if err != nil {
		return nil, processed_images.MakeErrorAddingImage(err)
	} //esto guarda la imagen en sí

	id, err := s.imagesDao.AddProcessedImageData(payload)
	if err != nil {
		return nil, processed_images.MakeErrorAddingImage(err)
	} //esto guarda el resultado de la imagen procesada

	//everything went ok
	code := "200"
	description := fmt.Sprintf("Bytes written: %d, Id: %s", bytesWritten, id)
	result := &processed_images.GoaResult{Code: &code, Description: &description}
	return result, nil
}

func (s *processedimagessrvc) GetProcessedSatelliteImage(ctx context.Context, payload *processed_images.GetProcessedSatelliteImagePayload) (result *processed_images.GoaResult, err error) {
	return nil, nil

}
