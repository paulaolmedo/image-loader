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
