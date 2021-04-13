package models

// ProcessedSatelliteImage is the payload type of the Processed images service
// Load new processed satellite image method.
type ProcessedSatelliteImage struct {
	// File name of the processed image
	FileName              *string                `json:"file_name,omitempty"`
	GeographicInformation *GeographicInformation `json:"geographic_information,omitempty"`
	// When was the image taken
	DateTime          *string            `json:"date_time,omitempty"`
	NormalizedIndexes *NormalizedIndexes `json:"normalized_indexes,omitempty"`
}

type GeographicInformation struct {
	// Non-forgetable identifier
	TagName *string
	// Coordinates of the satellite image
	Coordinates map[string]float64
}

type NormalizedIndexes struct {
	// Normalized difference vegetation index
	Ndvi []float64
	// Normalized difference water index
	Ndwi []float64
}

// RawSatelliteImage is the payload type of the Raw images service Load new raw
// satellite image method.
type RawSatelliteImage struct {
	// The image identifier
	ID *string `bson:"_id"`
	// File name of the raw image
	FileName *string
}
