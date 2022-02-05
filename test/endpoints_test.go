package test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	data "image-loader/internal/models"
	server "image-loader/server"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const mongoURL = "mongodb://mongo"

var (
	testserver        *httptest.Server
	testConfiguration server.Server
)

func TestMain(test *testing.M) {
	InitService()
	os.Exit(test.Run())
}

func InitService() {
	testConfiguration = server.Server{}
	testConfiguration.InitHTTPServer(mongoURL)
	testserver = httptest.NewServer(testConfiguration.Router)
}

func InitRawImage(Filename string) data.RawSatelliteImage {
	return data.RawSatelliteImage{Filename: Filename}
}

func InitProcessedImage(imageFilename string, resultsFileName string) data.ProcessedSatelliteImage {
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

	return data.ProcessedSatelliteImage{
		ImageFilename:         imageFilename,
		ResultsFilename:       resultsFileName,
		GeographicInformation: geoData,
		NormalizedIndexes:     normIndexes,
	}
}

func TestLoadRawImage(t *testing.T) {
	rawImage := InitRawImage("./test_image.tif")
	client := resty.New()

	response, err := client.R().
		SetBody(rawImage).
		EnableTrace().
		Post(testserver.URL + "/images/raw")

	require.NoError(t, err)

	// 4940427
	actualResponse := string(response.Body())
	expectedResponse := "[\"Bytes written while storing raw image: 4940427. \"]\n"

	actualStatusCode := response.StatusCode()
	expectedStatusCode := http.StatusCreated

	assert.Equal(t, expectedResponse, actualResponse)
	assert.Equal(t, actualStatusCode, expectedStatusCode)
	t.Logf("image correctly stored")
}

func TestGetRawImage(t *testing.T) {
	client := resty.New()

	response, err := client.R().
		SetQueryParams(map[string]string{
			"filename": "./test_image.tif",
		}).
		EnableTrace().
		Get(testserver.URL + "/images/raw")

	require.NoError(t, err)

	actualResponse := string(response.Body())
	expectedResponse := "[\"Bytes read: 4940427. \"]\n"

	actualStatusCode := response.StatusCode()
	expectedStatusCode := http.StatusOK

	assert.Equal(t, expectedResponse, actualResponse)
	assert.Equal(t, expectedStatusCode, actualStatusCode)
}

func TestGetNonExistentRawImage(t *testing.T) {
	client := resty.New()

	response, err := client.R().
		SetQueryParams(map[string]string{
			"filename": "./non_existent_file.tif",
		}).
		EnableTrace().
		Get(testserver.URL + "/images/raw")

	require.NoError(t, err)

	actualResponse := string(response.Body())
	expectedResponse := "[\"error retrieving image: file with given parameters not found\"]\n"

	actualStatusCode := response.StatusCode()
	expectedStatusCode := http.StatusInternalServerError

	assert.Equal(t, expectedResponse, actualResponse)
	assert.Equal(t, expectedStatusCode, actualStatusCode)
}

func TestLoadNonExistentRawImage(t *testing.T) {
	rawImage := InitRawImage("./non_existent_file.tif")
	client := resty.New()

	response, err := client.R().
		SetBody(rawImage).
		EnableTrace().
		Post(testserver.URL + "/images/raw")

	require.NoError(t, err)

	actualStatusCode := response.StatusCode()
	expectedStatusCode := http.StatusBadRequest

	assert.Equal(t, expectedStatusCode, actualStatusCode)
	t.Logf("no such image! error reading data")
}

func TestLoadErronousRawImage(t *testing.T) {
	rawImage := InitRawImage("./non_existent_file.someotherextension")
	client := resty.New()

	response, err := client.R().
		SetBody(rawImage).
		EnableTrace().
		Post(testserver.URL + "/images/raw")

	require.NoError(t, err)

	actualStatusCode := response.StatusCode()
	expectedStatusCode := http.StatusConflict

	assert.Equal(t, actualStatusCode, expectedStatusCode)
	t.Logf("did not recognized file extension")
}

func TestGetErronousRawImage(t *testing.T) {
	client := resty.New()

	response, err := client.R().
		SetQueryParams(map[string]string{
			"filename": "./non_existent_file.someotherextension",
		}).
		EnableTrace().
		Get(testserver.URL + "/images/raw")

	require.NoError(t, err)

	actualResponse := string(response.Body())
	expectedResponse := "[\"Unsuported filename.\"]\n"

	actualStatusCode := response.StatusCode()
	expectedStatusCode := http.StatusConflict

	assert.Equal(t, expectedResponse, actualResponse)
	assert.Equal(t, expectedStatusCode, actualStatusCode)
}

func TestLoadRawImageWithEmptyBody(t *testing.T) {
	client := resty.New()

	response, err := client.R().
		SetBody(nil).
		EnableTrace().
		Post(testserver.URL + "/images/raw")

	require.NoError(t, err)

	actualStatusCode := response.StatusCode()
	expectedStatusCode := http.StatusBadRequest

	assert.Equal(t, expectedStatusCode, actualStatusCode)
	t.Logf("invalid body")
}

func TestLoadProcessedImage(t *testing.T) {
	rawImage := InitProcessedImage("./test_image.tif", "./processing_results.csv")
	client := resty.New()

	response, err := client.R().
		SetBody(rawImage).
		EnableTrace().
		Post(testserver.URL + "/images/processed")

	require.NoError(t, err)

	actualStatusCode := response.StatusCode()
	expectedStatusCode := http.StatusCreated

	assert.Equal(t, expectedStatusCode, actualStatusCode)
	t.Logf("image correctly stored")
}

func TestLoadProcessedImageWithEmptyBody(t *testing.T) {
	client := resty.New()

	response, err := client.R().
		SetBody(nil).
		EnableTrace().
		Post(testserver.URL + "/images/processed")

	require.NoError(t, err)

	actualStatusCode := response.StatusCode()
	expectedStatusCode := http.StatusBadRequest

	assert.Equal(t, expectedStatusCode, actualStatusCode)
	t.Logf("invalid body")
}

func TestLoadNonExistentProcessedImage(t *testing.T) {
	rawImage := InitProcessedImage("./non_existent_file.tif", "")
	client := resty.New()

	response, err := client.R().
		SetBody(rawImage).
		EnableTrace().
		Post(testserver.URL + "/images/processed")

	require.NoError(t, err)

	actualStatusCode := response.StatusCode()
	expectedStatusCode := http.StatusBadRequest

	assert.Equal(t, expectedStatusCode, actualStatusCode)
	t.Logf("no such image! error reading data")
}

func TestGetInvalidProcessedImage(t *testing.T) {
	client := resty.New()

	response, err := client.R().
		SetQueryParams(map[string]string{
			"filename": "./processing_results.someotherextension",
		}).
		EnableTrace().
		Get(testserver.URL + "/images/processed")

	require.NoError(t, err)

	actualResponse := string(response.Body())
	expectedResponse := "[\"Unsuported filename.\"]\n"

	actualStatusCode := response.StatusCode()
	expectedStatusCode := http.StatusConflict

	assert.Equal(t, expectedResponse, actualResponse)
	assert.Equal(t, expectedStatusCode, actualStatusCode)
}
