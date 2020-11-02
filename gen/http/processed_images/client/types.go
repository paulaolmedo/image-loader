// Code generated by goa v3.2.3, DO NOT EDIT.
//
// Processed images HTTP client types
//
// Command:
// $ goa gen image-loader/design

package client

import (
	processedimages "image-loader/gen/processed_images"
	processedimagesviews "image-loader/gen/processed_images/views"

	goa "goa.design/goa/v3/pkg"
)

// LoadNewProcessedSatelliteImageRequestBody is the type of the "Processed
// images" service "Load new processed satellite image" endpoint HTTP request
// body.
type LoadNewProcessedSatelliteImageRequestBody struct {
	// The image identifier
	ID *string `bson:"_id"`
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
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// LoadNewProcessedSatelliteImageInternalErrorResponseBody is the type of the
// "Processed images" service "Load new processed satellite image" endpoint
// HTTP response body for the "InternalError" error.
type LoadNewProcessedSatelliteImageInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// GetProcessedSatelliteImageBadRequestResponseBody is the type of the
// "Processed images" service "Get processed satellite image" endpoint HTTP
// response body for the "BadRequest" error.
type GetProcessedSatelliteImageBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// GetProcessedSatelliteImageInternalErrorResponseBody is the type of the
// "Processed images" service "Get processed satellite image" endpoint HTTP
// response body for the "InternalError" error.
type GetProcessedSatelliteImageInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
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

// NewLoadNewProcessedSatelliteImageRequestBody builds the HTTP request body
// from the payload of the "Load new processed satellite image" endpoint of the
// "Processed images" service.
func NewLoadNewProcessedSatelliteImageRequestBody(p *processedimages.ProcessedSatelliteImage) *LoadNewProcessedSatelliteImageRequestBody {
	body := &LoadNewProcessedSatelliteImageRequestBody{
		ID:       p.ID,
		FileName: p.FileName,
		DateTime: p.DateTime,
	}
	if p.GeographicInformation != nil {
		body.GeographicInformation = marshalProcessedimagesGeographicInformationToGeographicInformationRequestBody(p.GeographicInformation)
	}
	if p.NormalizedIndexes != nil {
		body.NormalizedIndexes = marshalProcessedimagesNormalizedIndexesToNormalizedIndexesRequestBody(p.NormalizedIndexes)
	}
	return body
}

// NewGetProcessedSatelliteImageRequestBody builds the HTTP request body from
// the payload of the "Get processed satellite image" endpoint of the
// "Processed images" service.
func NewGetProcessedSatelliteImageRequestBody(p *processedimages.GetProcessedSatelliteImagePayload) *GetProcessedSatelliteImageRequestBody {
	body := &GetProcessedSatelliteImageRequestBody{
		FileName: p.FileName,
	}
	return body
}

// NewLoadNewProcessedSatelliteImageGoaResultOK builds a "Processed images"
// service "Load new processed satellite image" endpoint result from a HTTP
// "OK" response.
func NewLoadNewProcessedSatelliteImageGoaResultOK(body *LoadNewProcessedSatelliteImageResponseBody) *processedimagesviews.GoaResultView {
	v := &processedimagesviews.GoaResultView{
		Code:        body.Code,
		Description: body.Description,
	}

	return v
}

// NewLoadNewProcessedSatelliteImageBadRequest builds a Processed images
// service Load new processed satellite image endpoint BadRequest error.
func NewLoadNewProcessedSatelliteImageBadRequest(body *LoadNewProcessedSatelliteImageBadRequestResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewLoadNewProcessedSatelliteImageInternalError builds a Processed images
// service Load new processed satellite image endpoint InternalError error.
func NewLoadNewProcessedSatelliteImageInternalError(body *LoadNewProcessedSatelliteImageInternalErrorResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewGetProcessedSatelliteImageGoaResultOK builds a "Processed images" service
// "Get processed satellite image" endpoint result from a HTTP "OK" response.
func NewGetProcessedSatelliteImageGoaResultOK(body *GetProcessedSatelliteImageResponseBody) *processedimagesviews.GoaResultView {
	v := &processedimagesviews.GoaResultView{
		Code:        body.Code,
		Description: body.Description,
	}

	return v
}

// NewGetProcessedSatelliteImageBadRequest builds a Processed images service
// Get processed satellite image endpoint BadRequest error.
func NewGetProcessedSatelliteImageBadRequest(body *GetProcessedSatelliteImageBadRequestResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewGetProcessedSatelliteImageInternalError builds a Processed images service
// Get processed satellite image endpoint InternalError error.
func NewGetProcessedSatelliteImageInternalError(body *GetProcessedSatelliteImageInternalErrorResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// ValidateLoadNewProcessedSatelliteImageBadRequestResponseBody runs the
// validations defined on Load new processed satellite
// image_BadRequest_Response_Body
func ValidateLoadNewProcessedSatelliteImageBadRequestResponseBody(body *LoadNewProcessedSatelliteImageBadRequestResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateLoadNewProcessedSatelliteImageInternalErrorResponseBody runs the
// validations defined on Load new processed satellite
// image_InternalError_Response_Body
func ValidateLoadNewProcessedSatelliteImageInternalErrorResponseBody(body *LoadNewProcessedSatelliteImageInternalErrorResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateGetProcessedSatelliteImageBadRequestResponseBody runs the
// validations defined on Get processed satellite image_BadRequest_Response_Body
func ValidateGetProcessedSatelliteImageBadRequestResponseBody(body *GetProcessedSatelliteImageBadRequestResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateGetProcessedSatelliteImageInternalErrorResponseBody runs the
// validations defined on Get processed satellite
// image_InternalError_Response_Body
func ValidateGetProcessedSatelliteImageInternalErrorResponseBody(body *GetProcessedSatelliteImageInternalErrorResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}
