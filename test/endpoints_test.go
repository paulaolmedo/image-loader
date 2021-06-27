package test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	data "image-loader/models"
	server "image-loader/server"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

const mongoURL = "mongodb://localhost:27017"

var testserver *httptest.Server
var testConfiguration server.Server

func TestMain(test *testing.M) {
	InitService()
	os.Exit(test.Run())
}

func InitService() {
	testConfiguration = server.Server{}
	testConfiguration.InitHTTPServer(mongoURL)
	testserver = httptest.NewServer(testConfiguration.Router)
}

func InitRawImage(filename string) data.RawSatelliteImage {
	return data.RawSatelliteImage{FileName: filename}
}

func InitProcessedImage(filename string) data.ProcessedSatelliteImage {
	coordinates := map[string]float64{
		"x_min": -1.0,
		"y_min": -1.0,
		"x_max": 1.0,
		"y_max": 1.0,
	}
	geoData := data.GeographicInformation{TagName: "test-tag", Coordinates: coordinates}

	ndvi := []float64{1.0, 1.0}
	ndwi := []float64{1.0, 1.0}
	normIndexes := data.NormalizedIndexes{Ndvi: ndvi, Ndwi: ndwi}

	return data.ProcessedSatelliteImage{FileName: filename,
		GeographicInformation: geoData,
		NormalizedIndexes:     normIndexes}
}
func TestLoadRawImage(test *testing.T) {
	rawImage := InitRawImage("./test_image.tif")
	client := resty.New()

	response, err := client.R().
		SetBody(rawImage).
		EnableTrace().
		Post(testserver.URL + "/images/raw")

	if err != nil {
		test.Errorf("Error while calling endpoint: %v", err.Error())
	}
	//4940427
	data := string(response.Body())
	if err != nil {
		test.Errorf("Error converting data: %v", err.Error())
	}

	assert.Equal(test, data, "\"Bytes written 4940427\"\n")
	assert.Equal(test, response.StatusCode(), http.StatusCreated)
	test.Logf("image correctly stored")

}

func TestLoadNonExistentRawImage(test *testing.T) {
	rawImage := InitRawImage("./non_existent_file.tif")
	client := resty.New()

	response, err := client.R().
		SetBody(rawImage).
		EnableTrace().
		Post(testserver.URL + "/images/raw")

	if err != nil {
		test.Errorf("Error while calling endpoint: %v", err.Error())
	}

	assert.Equal(test, http.StatusBadRequest, response.StatusCode())
	test.Logf("no such image! error reading data")

}

func TestLoadProcessedImage(test *testing.T) {
	rawImage := InitProcessedImage("./test_image.tif")
	client := resty.New()

	response, err := client.R().
		SetBody(rawImage).
		EnableTrace().
		Post(testserver.URL + "/images/processed")

	if err != nil {
		test.Errorf("Error while calling endpoint: %v", err.Error())
	}

	assert.Equal(test, response.StatusCode(), http.StatusCreated)
	test.Logf("image correctly stored")

}

func TestLoadNonExistentProcessedImage(test *testing.T) {
	rawImage := InitProcessedImage("./non_existent_file.tif")
	client := resty.New()

	response, err := client.R().
		SetBody(rawImage).
		EnableTrace().
		Post(testserver.URL + "/images/processed")

	if err != nil {
		test.Errorf("Error while calling endpoint: %v", err.Error())
	}

	assert.Equal(test, http.StatusBadRequest, response.StatusCode())
	test.Logf("no such image! error reading data")

}
