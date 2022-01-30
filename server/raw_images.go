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
	data "image-loader/internal/models"
	"io/ioutil"
	"net/http"
	"strings"
)

func (s *Server) LoadNewRawSatelliteImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, appJSON)
	var imageProperties data.RawSatelliteImage

	err := json.NewDecoder(r.Body).Decode(&imageProperties)
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, errorReadingJSON)
		return
	}

	// le pongo .tif por ahora pero no sé qué tipos de imagen podemos guardar
	// TODO agregar algún otro tipo de validación para asegurarse que es un archivo "bueno"
	if !strings.Contains(imageProperties.Filename, tif) {
		jsonResponse(w, http.StatusConflict, errorReadingFilename)
		return
	}

	originalFile, err := ioutil.ReadFile(imageProperties.Filename)
	if err != nil {
		description := fmt.Sprintf(errorReadingFileData, err)
		jsonResponse(w, http.StatusBadRequest, description)
		return
	}

	response, err := s.Database.AddImage(originalFile, imageProperties.Filename, "raw")
	if err != nil {
		e := fmt.Errorf(errorStoringImage, err)
		jsonResponse(w, http.StatusInternalServerError, e)
		return
	} // guardo la imagen en sí

	jsonResponse(w, http.StatusCreated, response)
}

// Esto lo tendría que pasar por query params y no por el body
func (s *Server) GetRawSatelliteImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, appJSON)

	queryParams := r.URL.Query()
	filename := queryParams.Get("filename")

	if !strings.Contains(filename, tif) {
		jsonResponse(w, http.StatusConflict, errorReadingFilename)
		return
	}

	bytesRead, err := s.Database.GetImage(filename, "raw")
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	description := fmt.Sprintf(bRead, bytesRead)
	jsonResponse(w, http.StatusOK, description)
}
