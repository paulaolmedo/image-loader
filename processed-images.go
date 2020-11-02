package imageLoaderAPI

import (
	"context"
	processed_images "image-loader/gen/processed_images"
	"image-loader/mongo"
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
	return nil, nil
}

func (s *processedimagessrvc) GetProcessedSatelliteImage(ctx context.Context, payload *processed_images.GetProcessedSatelliteImagePayload) (result *processed_images.GoaResult, err error) {
	return nil, nil

}
