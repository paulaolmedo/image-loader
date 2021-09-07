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

package server

import (
	"encoding/json"
	"fmt"
	data "image-loader/models"
	"io/ioutil"
	"net/http"
	"strings"
)

//LoadNewProcessedSatelliteImage stores a new processed image
func (s *Server) LoadNewProcessedSatelliteImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, appJSON)

	var imageProperties data.ProcessedSatelliteImage

	err := json.NewDecoder(r.Body).Decode(&imageProperties)
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, errorReadingJSON)
		return
	}

	originalFile, err := ioutil.ReadFile(imageProperties.Filename)
	if err != nil {
		description := fmt.Sprintf(errorReadingFileData, err)
		jsonResponse(w, http.StatusBadRequest, description)
		return
	}

	bytesWritten, err := s.Database.AddImage(originalFile, imageProperties.Filename, "processed")
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, errorStoringImage)
		return
	}

	id, err := s.Database.AddProcessedImageData(&imageProperties)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, errorStoringDataImage)
		return
	}
	response := bWritten + "Id: %v"
	description := fmt.Sprintf(response, bytesWritten, id)
	jsonResponse(w, http.StatusCreated, description)
}

func (s *Server) GetProcessedSatelliteImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, appJSON)

	queryParams := r.URL.Query()
	filename := queryParams.Get("filename")

	// TODO agregar algún otro tipo de validación para asegurarse que es un archivo "bueno"
	if !strings.Contains(filename, csv) {
		jsonResponse(w, http.StatusConflict, errorReadingFilename)
		return
	}

	bytesRead, err := s.Database.GetImage(filename, "processed")
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, "Error retrieving processed image")
		return
	}

	description := fmt.Sprintf(bWritten, bytesRead)
	jsonResponse(w, http.StatusOK, description)
}
