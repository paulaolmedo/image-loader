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
	"github.com/stretchr/testify/require"
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

func InitRawImage(Filename string) data.RawSatelliteImage {
	return data.RawSatelliteImage{Filename: Filename}
}

func InitProcessedImage(Filename string) data.ProcessedSatelliteImage {
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

	return data.ProcessedSatelliteImage{Filename: Filename,
		GeographicInformation: geoData,
		NormalizedIndexes:     normIndexes}
}
func TestLoadRawImage(t *testing.T) {
	rawImage := InitRawImage("./test_image.tif")
	client := resty.New()

	response, err := client.R().
		SetBody(rawImage).
		EnableTrace().
		Post(testserver.URL + "/images/raw")

	require.NoError(t, err)

	//4940427
	actualResponse := string(response.Body())
	expectedResponse := "\"Bytes written: 4940427. \"\n"

	assert.Equal(t, expectedResponse, actualResponse)
	assert.Equal(t, response.StatusCode(), http.StatusCreated)
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
	expectedResponse := "\"Bytes read: 4940427. \"\n"

	assert.Equal(t, expectedResponse, actualResponse)
	assert.Equal(t, http.StatusOK, response.StatusCode())
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
	// TODO cuando no encuentra la imagen debería devolver 404
	expectedResponse := "\"Error retrieving raw image\"\n"

	assert.Equal(t, expectedResponse, actualResponse)
	assert.Equal(t, http.StatusInternalServerError, response.StatusCode())
}
func TestLoadNonExistentRawImage(t *testing.T) {
	rawImage := InitRawImage("./non_existent_file.tif")
	client := resty.New()

	response, err := client.R().
		SetBody(rawImage).
		EnableTrace().
		Post(testserver.URL + "/images/raw")

	require.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode())
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

	assert.Equal(t, http.StatusConflict, response.StatusCode())
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
	expectedResponse := "\"Unsuported filename.\"\n"

	assert.Equal(t, expectedResponse, actualResponse)
	assert.Equal(t, http.StatusConflict, response.StatusCode())
}

func TestLoadRawImageWithEmptyBody(t *testing.T) {
	client := resty.New()

	response, err := client.R().
		SetBody(nil).
		EnableTrace().
		Post(testserver.URL + "/images/raw")

	require.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode())
	t.Logf("invalid body")
}

func TestLoadProcessedImage(t *testing.T) {
	rawImage := InitProcessedImage("./test_image.tif")
	client := resty.New()

	response, err := client.R().
		SetBody(rawImage).
		EnableTrace().
		Post(testserver.URL + "/images/processed")

	require.NoError(t, err)

	assert.Equal(t, response.StatusCode(), http.StatusCreated)
	t.Logf("image correctly stored")
}

func TestLoadNonExistentProcessedImage(t *testing.T) {
	rawImage := InitProcessedImage("./non_existent_file.tif")
	client := resty.New()

	response, err := client.R().
		SetBody(rawImage).
		EnableTrace().
		Post(testserver.URL + "/images/processed")

	require.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode())
	t.Logf("no such image! error reading data")
}
