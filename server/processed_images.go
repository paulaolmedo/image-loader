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

// LoadNewProcessedSatelliteImage stores a new processed image
// TODO al iguala que las imágenes no procesadas, controlar el contenido de lo que se está guardando
// Carga tanto las imágenes procesadas como los resultados de procesamiento
func (server *Server) LoadNewProcessedSatelliteImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, appJSON)

	var imageProperties data.ProcessedSatelliteImage

	err := json.NewDecoder(r.Body).Decode(&imageProperties)
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, errorReadingJSON)
		return
	}

	imageFilename := imageProperties.ImageFilename
	resultsFilename := imageProperties.ResultsFilename

	// TODO Mejorar estas validaciones algún día
	image, errorDescription := checkIfFileIsValid("png", imageFilename)
	if errorDescription != "" {
		jsonResponse(w, http.StatusBadRequest, errorDescription)
		return
	}

	// revisar esto porque los resultados siempre deberían estar
	results, errorDescription := checkIfFileIsValid(csv, resultsFilename)
	if errorDescription != "" {
		jsonResponse(w, http.StatusBadRequest, errorDescription)
		return
	}

	// primero va a agregar la imagen en caso de que exista
	var responseFromImage string
	if image != nil {
		responseFromImage, err = server.Database.AddImage(image, imageFilename, "processed")
		if err != nil {
			jsonResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	// agrega los resultados del procesamiento (esto sí que tiene que hacerse si o si)
	responseFromResults, err := server.Database.AddImage(results, resultsFilename, "results", &imageProperties)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := responseFromImage + responseFromResults
	jsonResponse(w, http.StatusCreated, response)
}

// GetProcessedSatelliteImage devuelve la imagen .png + los resultados de procesamiento, para pasárselos al otro servicio
func (server *Server) GetProcessedSatelliteImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, "image/png")

	queryParams := r.URL.Query()
	filename := queryParams.Get("filename")

	// TODO agregar algún otro tipo de validación para asegurarse que es un archivo "bueno"
	if !strings.Contains(filename, ".png") {
		jsonResponse(w, http.StatusConflict, errorReadingFilename)
		return
	}

	_, err := server.Database.GetImage(filename, "processed")
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, err)
		return
	}
 
	http.ServeFile(w, r, filename)
}

func checkIfFileIsValid(extension string, filename string) ([]byte, string) {
	var image []byte
	var err error
	if filename != "" && strings.Contains(filename, extension) {
		image, err = ioutil.ReadFile(filename)
		if err != nil {
			return nil, fmt.Sprintf(errorReadingFileData, err)
		}
	}

	return image, ""
}
