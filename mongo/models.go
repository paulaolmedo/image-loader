package mongo

import (
	images "image-loader/gen/images"
	"time"
)

const (
	//images  = "satellite_images"
	idField = "_id"
)

//SatelliteImageDetails represents the satellite image itself
type SatelliteImageDetails struct {
	ID                           string `bson:"_id"`
	images.GeographicInformation `bson:"geographic_information"`
	images.NormalizedIndexes     `bson:"normalized_indexes"`
	dateTime                     time.Time `bson:"date_time"`
}
