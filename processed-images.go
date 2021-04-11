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
