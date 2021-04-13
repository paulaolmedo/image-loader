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
	"errors"
	processed_images "image-loader/gen/processed_images"
)

//properties of the database service
type dbProperties struct {
	imageRepository ImageRepository
}

//DatabaseService contiene las funciones para manipular la bd
type DatabaseService interface {
	AddImage(originalFile []byte, filename string, imageType string) (int, error)
	GetImage(filename string, imageType string) (int64, error)
	AddProcessedImageData(image *processed_images.ProcessedSatelliteImage) (string, error)
}

//NewImageService inicializa el servicio de basw de datos
func NewImageService(imageRepository ImageRepository) DatabaseService {
	return &dbProperties{imageRepository}
}

func (properties *dbProperties) AddImage(originalFile []byte, filename string, imageType string) (int, error) {
	if imageType == "raw" {
		size, err := properties.imageRepository.AddRawImage(originalFile, filename)
		if err != nil {
			return 0, errors.New("an error ocurred while storing the raw image")
		}
		return size, nil
	} else if imageType == "processed" {
		size, err := properties.imageRepository.AddProcessedImage(originalFile, filename)
		if err != nil {
			return 0, errors.New("an error ocurred while storing the raw image")
		}
		return size, nil
	} else {
		return 0, errors.New("wrong image type selected")
	}
}

func (properties *dbProperties) AddProcessedImageData(image *processed_images.ProcessedSatelliteImage) (string, error) {
	result, err := properties.imageRepository.AddProcessedImageData(image)
	if err != nil {
		return "", errors.New("an error ocurred while storing the processed image data")
	}

	return result, nil
}

func (properties *dbProperties) GetImage(filename string, imageType string) (int64, error) {
	if imageType == "raw" {
		size, err := properties.imageRepository.GetRawImage(filename)
		if err != nil {
			return 0, errors.New("an error ocurred while retrieving the raw image")
		}
		return size, nil
	} else if imageType == "processed" {
		size, err := properties.imageRepository.GetProcessedImage(filename)
		if err != nil {
			return 0, errors.New("an error ocurred while retrieving the raw image")
		}
		return size, nil
	} else {
		return 0, errors.New("wrong image type selected")
	}
}