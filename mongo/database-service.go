package mongo

import processed_images "image-loader/gen/processed_images"

//properties of the database service
type dbProperties struct {
	imageRepository ImageRepository
}

//DatabaseService contiene las funciones para manipular la bd
type DatabaseService interface {
	AddRawImage(originalFile []byte, filename string) (int, error)
	AddProcessedImage(originalFile []byte, filename string) (int, error)
	AddProcessedImageData(image *processed_images.ProcessedSatelliteImage) (string, error)
	GetRawImage(filename string) (int64, error)
	GetProcessedImage(filename string) (int64, error)
}

//NewImageService inicializa el servicio de basw de datos
func NewImageService(imageRepository ImageRepository) DatabaseService {
	return &dbProperties{imageRepository}
}

func (properties *dbProperties) AddRawImage(originalFile []byte, filename string) (int, error) {
	return 0, nil
}

func (properties *dbProperties) AddProcessedImage(originalFile []byte, filename string) (int, error) {
	return 0, nil
}

func (properties *dbProperties) AddProcessedImageData(image *processed_images.ProcessedSatelliteImage) (string, error) {
	return "", nil

}

func (properties *dbProperties) GetRawImage(filename string) (int64, error) {
	return 0, nil
}

func (properties *dbProperties) GetProcessedImage(filename string) (int64, error) {
	return 0, nil
}
