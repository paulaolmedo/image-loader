package imageLoaderAPI

import (
	"context"
	images "image-loader/gen/images"
	"image-loader/mongo"
	"log"
)

type imagesrvc struct {
	logger          *log.Logger
	imageRepository mongo.ImageRepository
}

//NewImageLoader .
func NewImageLoader(logger *log.Logger, imageRepository mongo.ImageRepository) images.Service {
	return &imagesrvc{logger, imageRepository}
}

func (s *imagesrvc) LoadNewSatelliteImage(ctx context.Context, payload *images.SatelliteImage) (err error) {
	return nil
}
