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

// RawSatelliteImage Information (raw)
var RawSatelliteImage = Type("RawSatelliteImage", func() {
	Attribute("id", String, "The image identifier", func() {
		Meta("struct:tag:bson", "_id")
	})
	Attribute("file_name", String, "File name of the raw image")
})

//ProcessedSatelliteImage Information (not raw, clearly)
var ProcessedSatelliteImage = Type("ProcessedSatelliteImage", func() {
	Attribute("file_name", String, "File name of the processed image")
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

// OperationResult of an operation
var OperationResult = ResultType("vnd.application/goa.Result+json", func() {
	Description("The results of an operation")
	Attributes(func() {
		Attribute("code", String, "The operation code")
		Attribute("description", String, "The operation description")
	})
	View("default", func() {
		Attribute("code")
		Attribute("description")
	})
})

/********** Services definition **********/
var _ = Service("Raw images", func() {
	HTTP(func() {
		Path("/raw-images")
	})
	Error("InternalError")
	Error("BadRequest")
	Error("NotFound")
	Method("Load new raw satellite image", func() {
		Description("loads a new image into the database")
		Payload(RawSatelliteImage)
		Result(OperationResult)
		Error("ErrorAddingImage")
		HTTP(func() {
			POST("/")
			Response(StatusOK)
			Response("BadRequest", StatusBadRequest)
			Response("InternalError", StatusInternalServerError)
		})
	})
	Method("Get raw satellite image", func() {
		Description("get a raw image from the database")
		Payload(func() {
			Attribute("file_name", String, "File name of the raw image")
		})
		Result(OperationResult)
		Error("ErrorGettingImage")
		HTTP(func() {
			GET("/")
			Response(StatusOK)
			Response("BadRequest", StatusBadRequest)
			Response("InternalError", StatusInternalServerError)
		})
	})
})

var _ = Service("Processed images", func() {
	HTTP(func() {
		Path("/processed-images")
	})
	Error("InternalError")
	Error("BadRequest")
	Error("NotFound")
	Method("Load new processed satellite image", func() {
		Description("loads a new processed image into the database")
		Payload(ProcessedSatelliteImage)
		Result(OperationResult)
		Error("ErrorAddingImage")
		HTTP(func() {
			POST("/")
			Response(StatusOK)
			Response("BadRequest", StatusBadRequest)
			Response("InternalError", StatusInternalServerError)
		})
	})
	Method("Get processed satellite image", func() {
		Description("get a processed image from the database")
		Payload(func() {
			Attribute("file_name", String, "File name of the processed image")
		})
		Result(OperationResult)
		Error("ErrorGettingImage")
		HTTP(func() {
			GET("/")
			Response(StatusOK)
			Response("BadRequest", StatusBadRequest)
			Response("InternalError", StatusInternalServerError)
		})
	})
})
