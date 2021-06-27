package models

import "time"

// ProcessedSatelliteImage
//
// Payload type of the processed images service
//
// swagger:model
type ProcessedSatelliteImage struct {
	// Filename of the processing results (.csv)
	Filename string `json:"file_name,omitempty"`
	// General geographic information associated to the image
	GeographicInformation GeographicInformation `json:"geographic_information,omitempty"`
	// When was the image taken
	DateTime time.Time `json:"date_time,omitempty"`
	// @wip definir si es necesario o si me voy a quedar con la información del archivo nada más
	NormalizedIndexes NormalizedIndexes `json:"normalized_indexes,omitempty"`
}

// RawSatelliteImage
//
// Payload type of the raw images service
//
// swagger:model
type RawSatelliteImage struct {
	// The image identifier
	ID *string `bson:"_id"`
	// Filename of the raw image
	Filename string `json:"file_name,omitempty"`
}

// GeographicInformation
//
// Geographic information of an image
//
// swagger:model
type GeographicInformation struct {
	// Non-forgetable identifier
	TagName string
	// Main coordinates of the processed image
	Coordinates map[string]float64
}

// NormalizedIndexes
//
// @wip
//
// swagger:model
type NormalizedIndexes struct {
	// Normalized difference vegetation index
	Ndvi []float64
	// Normalized difference water index
	Ndwi []float64
}

//ModelError
//
// @wip
//
// swagger:model
type ModelError struct {
	//código de error
	Code string `json:"code"`
	//mensaje
	Message string `json:"message"`
}
