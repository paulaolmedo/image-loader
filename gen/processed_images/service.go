// Code generated by goa v3.3.1, DO NOT EDIT.
//
// Processed images service
//
// Command:
// $ goa gen image-loader/design

package processedimages

import (
	"context"
	processedimagesviews "image-loader/gen/processed_images/views"

	goa "goa.design/goa/v3/pkg"
)

// Service is the Processed images service interface.
type Service interface {
	// loads a new processed image into the database
	LoadNewProcessedSatelliteImage(context.Context, *ProcessedSatelliteImage) (res *GoaResult, err error)
	// get a processed image from the database
	GetProcessedSatelliteImage(context.Context, *GetProcessedSatelliteImagePayload) (res *GoaResult, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "Processed images"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [2]string{"Load new processed satellite image", "Get processed satellite image"}

// ProcessedSatelliteImage is the payload type of the Processed images service
// Load new processed satellite image method.
type ProcessedSatelliteImage struct {
	// File name of the processed image
	FileName              *string
	GeographicInformation *GeographicInformation
	// When was the image taken
	DateTime          *string
	NormalizedIndexes *NormalizedIndexes
}

// GoaResult is the result type of the Processed images service Load new
// processed satellite image method.
type GoaResult struct {
	// The operation code
	Code *string
	// The operation description
	Description *string
}

// GetProcessedSatelliteImagePayload is the payload type of the Processed
// images service Get processed satellite image method.
type GetProcessedSatelliteImagePayload struct {
	// File name of the processed image
	FileName *string
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

// MakeInternalError builds a goa.ServiceError from an error.
func MakeInternalError(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "InternalError",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeBadRequest builds a goa.ServiceError from an error.
func MakeBadRequest(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "BadRequest",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeNotFound builds a goa.ServiceError from an error.
func MakeNotFound(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "NotFound",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeErrorAddingImage builds a goa.ServiceError from an error.
func MakeErrorAddingImage(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "ErrorAddingImage",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeErrorGettingImage builds a goa.ServiceError from an error.
func MakeErrorGettingImage(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "ErrorGettingImage",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// NewGoaResult initializes result type GoaResult from viewed result type
// GoaResult.
func NewGoaResult(vres *processedimagesviews.GoaResult) *GoaResult {
	return newGoaResult(vres.Projected)
}

// NewViewedGoaResult initializes viewed result type GoaResult from result type
// GoaResult using the given view.
func NewViewedGoaResult(res *GoaResult, view string) *processedimagesviews.GoaResult {
	p := newGoaResultView(res)
	return &processedimagesviews.GoaResult{Projected: p, View: "default"}
}

// newGoaResult converts projected type GoaResult to service type GoaResult.
func newGoaResult(vres *processedimagesviews.GoaResultView) *GoaResult {
	res := &GoaResult{
		Code:        vres.Code,
		Description: vres.Description,
	}
	return res
}

// newGoaResultView projects result type GoaResult to projected type
// GoaResultView using the "default" view.
func newGoaResultView(res *GoaResult) *processedimagesviews.GoaResultView {
	vres := &processedimagesviews.GoaResultView{
		Code:        res.Code,
		Description: res.Description,
	}
	return vres
}
