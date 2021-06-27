package models

// ProcessedSatelliteImage Payload type of the processed images service
//
// swagger:model
type ProcessedSatelliteImage struct {
	// File name of the processed image
	FileName string `json:"file_name,omitempty"`
	// @wip description
	GeographicInformation GeographicInformation `json:"geographic_information,omitempty"`
	// When was the image taken
	DateTime *string `json:"date_time,omitempty"` // transformar a time.Time
	// @wip description
	NormalizedIndexes NormalizedIndexes `json:"normalized_indexes,omitempty"`
}

// GeographicInformation @wip description
//
// swagger:model
type GeographicInformation struct {
	// Non-forgetable identifier
	TagName string
	// Coordinates of the satellite image
	Coordinates map[string]float64
}

// NormalizedIndexes @wip description
//
// swagger:model
type NormalizedIndexes struct {
	// Normalized difference vegetation index
	Ndvi []float64
	// Normalized difference water index
	Ndwi []float64
}

// RawSatelliteImage Payload type of the raw images service
//
// swagger:model
type RawSatelliteImage struct {
	// The image identifier
	ID *string `bson:"_id"`
	// File name of the raw image
	FileName string `json:"file_name,omitempty"`
}

//ModelError .
//
// swagger:model
type ModelError struct {
	//código de error
	Code string `json:"code"`
	//mensaje
	Message string `json:"message"`
}
