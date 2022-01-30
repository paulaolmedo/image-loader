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
	"fmt"
	processed_images "image-loader/internal/models"
)

const (
	errorRetrievingImage = "error retrieving %v image: %v"
)

// properties of the database service
type dbProperties struct {
	imageRepository ImageRepository
}

// DatabaseService contiene las funciones para manipular la bd
type DatabaseService interface {
	AddImage(image []byte, imageFilename string, operation string, imageProperties ...*processed_images.ProcessedSatelliteImage) (string, error)
	GetImage(Filename string, imageType string) (int64, error)
}

// NewImageService inicializa el servicio de basw de datos
func NewImageService(imageRepository ImageRepository) DatabaseService {
	return &dbProperties{imageRepository}
}

func (properties *dbProperties) AddImage(image []byte, imageFilename string, operation string, imageProperties ...*processed_images.ProcessedSatelliteImage) (string, error) {
	size, err := properties.imageRepository.AddFile(image, imageFilename, operation)
	if err != nil {
		return "", err
	}

	response := "Bytes written while storing %v image: %v. "

	if operation == "results" {
		result, err := properties.imageRepository.AddProcessedImageData(imageProperties...)
		if err != nil {
			return "", err
		}

		response += "Id of the stored results: " + result
	}

	description := fmt.Sprintf(response, operation, size)
	return description, nil
}

func (properties *dbProperties) GetImage(filename string, imageType string) (int64, error) {
	var err error
	var size int64

	switch imageType {
	case "raw":
		size, err = properties.imageRepository.GetRawImage(filename)

	case "processed":
		size, err = properties.imageRepository.GetProcessedImage(filename)

	default:
		err = errors.New("wrong image type selected")
	}

	if err != nil {
		errorMessage := fmt.Errorf(errorRetrievingImage, imageType, err)
		return 0, errorMessage
	}

	return size, nil
}
