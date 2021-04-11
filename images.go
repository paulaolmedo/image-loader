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

func (s *Server) LoadNewRawSatelliteImage(w http.ResponseWriter, r *http.Request) {

	/*filename := "*payload.FileName"
	if filename == "" {
		return nil, raw_images.MakeBadRequest(errors.New("File name cannot be empty"))
	}

	originalFile, err := ioutil.ReadFile(filename)
	//error reading original file
	if err != nil {
		return nil, raw_images.MakeBadRequest(err)
	}

	bytesWritten, err := s.Database.AddRawImage(originalFile, filename)
	//error adding image to database
	if err != nil {
		return nil, raw_images.MakeErrorAddingImage(err)
	}

	//everything went ok
	code := "200"
	description := fmt.Sprintf("Bytes written %d", bytesWritten)
	result = &raw_images.GoaResult{Code: &code, Description: &description}
	return result, nil*/
}

func (s *Server) GetRawSatelliteImage(w http.ResponseWriter, r *http.Request) {
	/*filename := *payload.FileName
	if filename == "" {
		return nil, raw_images.MakeBadRequest(errors.New("File name cannot be empty"))
	}

	bytesRead, err := s.Database.GetRawImage(filename)

	if err != nil {
		return nil, raw_images.MakeErrorGettingImage(err)
	}

	code := "200"
	description := fmt.Sprintf("Bytes read %d", bytesRead)
	result = &raw_images.GoaResult{Code: &code, Description: &description}
	return result, nil*/
}
