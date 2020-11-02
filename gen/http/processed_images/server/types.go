// Code generated by goa v3.2.3, DO NOT EDIT.
//
// Processed images HTTP server types
//
// Command:
// $ goa gen image-loader/design

package server

import (
	processedimages "image-loader/gen/processed_images"
	processedimagesviews "image-loader/gen/processed_images/views"

	goa "goa.design/goa/v3/pkg"
)

// LoadNewProcessedSatelliteImageRequestBody is the type of the "Processed
// images" service "Load new processed satellite image" endpoint HTTP request
// body.
type LoadNewProcessedSatelliteImageRequestBody struct {
	// File name of the processed image
	FileName              *string                           `form:"file_name,omitempty" json:"file_name,omitempty" xml:"file_name,omitempty"`
	GeographicInformation *GeographicInformationRequestBody `form:"geographic_information,omitempty" json:"geographic_information,omitempty" xml:"geographic_information,omitempty"`
	// When was the image taken
	DateTime          *string                       `form:"date_time,omitempty" json:"date_time,omitempty" xml:"date_time,omitempty"`
	NormalizedIndexes *NormalizedIndexesRequestBody `form:"normalized_indexes,omitempty" json:"normalized_indexes,omitempty" xml:"normalized_indexes,omitempty"`
}

// GetProcessedSatelliteImageRequestBody is the type of the "Processed images"
// service "Get processed satellite image" endpoint HTTP request body.
type GetProcessedSatelliteImageRequestBody struct {
	// File name of the processed image
	FileName *string `form:"file_name,omitempty" json:"file_name,omitempty" xml:"file_name,omitempty"`
}

// LoadNewProcessedSatelliteImageResponseBody is the type of the "Processed
// images" service "Load new processed satellite image" endpoint HTTP response
// body.
type LoadNewProcessedSatelliteImageResponseBody struct {
	// The operation code
	Code *string `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// The operation description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
}

// GetProcessedSatelliteImageResponseBody is the type of the "Processed images"
// service "Get processed satellite image" endpoint HTTP response body.
type GetProcessedSatelliteImageResponseBody struct {
	// The operation code
	Code *string `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// The operation description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
}

// LoadNewProcessedSatelliteImageBadRequestResponseBody is the type of the
// "Processed images" service "Load new processed satellite image" endpoint
// HTTP response body for the "BadRequest" error.
type LoadNewProcessedSatelliteImageBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// LoadNewProcessedSatelliteImageInternalErrorResponseBody is the type of the
// "Processed images" service "Load new processed satellite image" endpoint
// HTTP response body for the "InternalError" error.
type LoadNewProcessedSatelliteImageInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// GetProcessedSatelliteImageBadRequestResponseBody is the type of the
// "Processed images" service "Get processed satellite image" endpoint HTTP
// response body for the "BadRequest" error.
type GetProcessedSatelliteImageBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// GetProcessedSatelliteImageInternalErrorResponseBody is the type of the
// "Processed images" service "Get processed satellite image" endpoint HTTP
// response body for the "InternalError" error.
type GetProcessedSatelliteImageInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// GeographicInformationRequestBody is used to define fields on request body
// types.
type GeographicInformationRequestBody struct {
	// Non-forgetable identifier
	TagName *string `form:"tag_name,omitempty" json:"tag_name,omitempty" xml:"tag_name,omitempty"`
	// Coordinates of the satellite image
	Coordinates map[string]float64 `form:"coordinates,omitempty" json:"coordinates,omitempty" xml:"coordinates,omitempty"`
}

// NormalizedIndexesRequestBody is used to define fields on request body types.
type NormalizedIndexesRequestBody struct {
	// Normalized difference vegetation index
	Ndvi []float64 `form:"ndvi,omitempty" json:"ndvi,omitempty" xml:"ndvi,omitempty"`
	// Normalized difference water index
	Ndwi []float64 `form:"ndwi,omitempty" json:"ndwi,omitempty" xml:"ndwi,omitempty"`
}

// NewLoadNewProcessedSatelliteImageResponseBody builds the HTTP response body
// from the result of the "Load new processed satellite image" endpoint of the
// "Processed images" service.
func NewLoadNewProcessedSatelliteImageResponseBody(res *processedimagesviews.GoaResultView) *LoadNewProcessedSatelliteImageResponseBody {
	body := &LoadNewProcessedSatelliteImageResponseBody{
		Code:        res.Code,
		Description: res.Description,
	}
	return body
}

// NewGetProcessedSatelliteImageResponseBody builds the HTTP response body from
// the result of the "Get processed satellite image" endpoint of the "Processed
// images" service.
func NewGetProcessedSatelliteImageResponseBody(res *processedimagesviews.GoaResultView) *GetProcessedSatelliteImageResponseBody {
	body := &GetProcessedSatelliteImageResponseBody{
		Code:        res.Code,
		Description: res.Description,
	}
	return body
}

// NewLoadNewProcessedSatelliteImageBadRequestResponseBody builds the HTTP
// response body from the result of the "Load new processed satellite image"
// endpoint of the "Processed images" service.
func NewLoadNewProcessedSatelliteImageBadRequestResponseBody(res *goa.ServiceError) *LoadNewProcessedSatelliteImageBadRequestResponseBody {
	body := &LoadNewProcessedSatelliteImageBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewLoadNewProcessedSatelliteImageInternalErrorResponseBody builds the HTTP
// response body from the result of the "Load new processed satellite image"
// endpoint of the "Processed images" service.
func NewLoadNewProcessedSatelliteImageInternalErrorResponseBody(res *goa.ServiceError) *LoadNewProcessedSatelliteImageInternalErrorResponseBody {
	body := &LoadNewProcessedSatelliteImageInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetProcessedSatelliteImageBadRequestResponseBody builds the HTTP response
// body from the result of the "Get processed satellite image" endpoint of the
// "Processed images" service.
func NewGetProcessedSatelliteImageBadRequestResponseBody(res *goa.ServiceError) *GetProcessedSatelliteImageBadRequestResponseBody {
	body := &GetProcessedSatelliteImageBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetProcessedSatelliteImageInternalErrorResponseBody builds the HTTP
// response body from the result of the "Get processed satellite image"
// endpoint of the "Processed images" service.
func NewGetProcessedSatelliteImageInternalErrorResponseBody(res *goa.ServiceError) *GetProcessedSatelliteImageInternalErrorResponseBody {
	body := &GetProcessedSatelliteImageInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewLoadNewProcessedSatelliteImageProcessedSatelliteImage builds a Processed
// images service Load new processed satellite image endpoint payload.
func NewLoadNewProcessedSatelliteImageProcessedSatelliteImage(body *LoadNewProcessedSatelliteImageRequestBody) *processedimages.ProcessedSatelliteImage {
	v := &processedimages.ProcessedSatelliteImage{
		FileName: body.FileName,
		DateTime: body.DateTime,
	}
	if body.GeographicInformation != nil {
		v.GeographicInformation = unmarshalGeographicInformationRequestBodyToProcessedimagesGeographicInformation(body.GeographicInformation)
	}
	if body.NormalizedIndexes != nil {
		v.NormalizedIndexes = unmarshalNormalizedIndexesRequestBodyToProcessedimagesNormalizedIndexes(body.NormalizedIndexes)
	}

	return v
}

// NewGetProcessedSatelliteImagePayload builds a Processed images service Get
// processed satellite image endpoint payload.
func NewGetProcessedSatelliteImagePayload(body *GetProcessedSatelliteImageRequestBody) *processedimages.GetProcessedSatelliteImagePayload {
	v := &processedimages.GetProcessedSatelliteImagePayload{
		FileName: body.FileName,
	}

	return v
}

// ValidateLoadNewProcessedSatelliteImageRequestBody runs the validations
// defined on Load New Processed Satellite ImageRequestBody
func ValidateLoadNewProcessedSatelliteImageRequestBody(body *LoadNewProcessedSatelliteImageRequestBody) (err error) {
	if body.DateTime != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.date_time", *body.DateTime, goa.FormatDateTime))
	}
	return
}
