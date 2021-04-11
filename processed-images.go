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

package imageloaderapi

import (
	"net/http"
)

func (s *Server) LoadNewProcessedSatelliteImage(w http.ResponseWriter, r *http.Request) {
	/*filename := *payload.FileName
	if filename == "" {
		return nil, processed_images.MakeBadRequest(errors.New("File name cannot be empty"))
	}

	originalFile, err := ioutil.ReadFile(filename)
	//error reading original file
	if err != nil {
		return nil, processed_images.MakeBadRequest(err)
	}

	bytesWritten, err := s.Database.AddProcessedImage(originalFile, filename)
	//error adding image to database
	if err != nil {
		return nil, processed_images.MakeErrorAddingImage(err)
	} //esto guarda la imagen en sí

	id, err := s.Database.AddProcessedImageData(payload)
	if err != nil {
		return nil, processed_images.MakeErrorAddingImage(err)
	} //esto guarda el resultado de la imagen procesada

	//everything went ok
	code := "200"
	description := fmt.Sprintf("Bytes written: %d, Id: %s", bytesWritten, id)
	result := &processed_images.GoaResult{Code: &code, Description: &description}
	return result, nil*/
}

func (s *Server) GetProcessedSatelliteImage(w http.ResponseWriter, r *http.Request) {

}
