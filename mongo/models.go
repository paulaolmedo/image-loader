package mongo

import (
	processed_images "image-loader/gen/processed_images"
	"time"
)

const (
	//images  = "satellite_images"
	idField = "_id"
	db      = "safo"
)

// SessionState type
type SessionState string

// SessionState values
const (
	Pending     SessionState = "PENDING"
	Active      SessionState = "ACTIVE"
	Blocked     SessionState = "BLOCKED"
	Invalidated SessionState = "INVALIDATED"
)

//RawSatelliteImageDetails represents the satellite image itself
type RawSatelliteImageDetails struct {
	ID       string    `bson:"_id"`
	fileName string    `bson:"file_name"`
	dateTime time.Time `bson:"date_time"`
}

//ProcessedSatelliteImageDetails represents the satellite image already processed by the python service
type ProcessedSatelliteImageDetails struct {
	ID                                     string    `bson:"_id"`
	dateTime                               time.Time `bson:"date_time"`
	processed_images.GeographicInformation `bson:"geographic_information"`
	processed_images.NormalizedIndexes     `bson:"normalized_indexes"`
}
