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
