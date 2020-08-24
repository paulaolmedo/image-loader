package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("Satellite Image Loader API", func() {
	Title("Satellite Image Loader")
	Description("Service to interact with satellite images on a database")
	Version("1.0")
	HTTP(func() {
		Path("/image-loader/api/v1")
	})
})

/********** Data type definition **********/

// SatelliteImage Information
var SatelliteImage = Type("SatelliteImage", func() {
	Attribute("id", String, "The image identifier", func() {
		Meta("struct:tag:bson", "_id")
	})
	Attribute("geographic_information", GeographicInformation)
	Attribute("date_time", String, "When was the image taken", func() {
		Format(FormatDateTime)
	})
	Attribute("normalized_indexes", NormalizedIndexes)
})

// GeographicInformation .
var GeographicInformation = Type("GeographicInformation", func() {
	Attribute("tag_name", String, "Non-forgetable identifier")
	Attribute("coordinates", MapOf(String, Float64), "Coordinates of the satellite image")
})

//NormalizedIndexes .
var NormalizedIndexes = Type("NormalizedIndexes", func() {
	Attribute("ndvi", ArrayOf(Float64), "Normalized difference vegetation index")
	Attribute("ndwi", ArrayOf(Float64), "Normalized difference water index")
	//@wip: here can be added: enhanced vegetation index, burning area index...
})

/********** Data type definition **********/
var _ = Service("Images", func() {
	HTTP(func() {
		Path("/images")
	})
	Error("InternalError")
	Error("BadRequest")
	Error("NotFound")
	Method("Load new satellite image", func() {
		Description("loads a new image into the database")
		Payload(SatelliteImage)
		Result(func() {
			Description("Image added")
		})
		Error("ErrorAddingImage")
		HTTP(func() {
			POST("/")
			Response(StatusOK)
			Response("BadRequest", StatusBadRequest)
			Response("InternalError", StatusInternalServerError)
		})
	})
})
