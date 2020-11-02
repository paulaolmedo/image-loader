// Code generated by goa v3.2.3, DO NOT EDIT.
//
// Raw images service
//
// Command:
// $ goa gen image-loader/design

package rawimages

import (
	"context"
	rawimagesviews "image-loader/gen/raw_images/views"

	goa "goa.design/goa/v3/pkg"
)

// Service is the Raw images service interface.
type Service interface {
	// loads a new image into the database
	LoadNewRawSatelliteImage(context.Context, *RawSatelliteImage) (res *GoaResult, err error)
	// get a raw image from the database
	GetRawSatelliteImage(context.Context, *GetRawSatelliteImagePayload) (res *GoaResult, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "Raw images"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [2]string{"Load new raw satellite image", "Get raw satellite image"}

// RawSatelliteImage is the payload type of the Raw images service Load new raw
// satellite image method.
type RawSatelliteImage struct {
	// The image identifier
	ID *string `bson:"_id"`
	// File name of the raw image
	FileName *string
}

// GoaResult is the result type of the Raw images service Load new raw
// satellite image method.
type GoaResult struct {
	// The operation code
	Code *string
	// The operation description
	Description *string
}

// GetRawSatelliteImagePayload is the payload type of the Raw images service
// Get raw satellite image method.
type GetRawSatelliteImagePayload struct {
	// File name of the raw image
	FileName *string
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
func NewGoaResult(vres *rawimagesviews.GoaResult) *GoaResult {
	return newGoaResult(vres.Projected)
}

// NewViewedGoaResult initializes viewed result type GoaResult from result type
// GoaResult using the given view.
func NewViewedGoaResult(res *GoaResult, view string) *rawimagesviews.GoaResult {
	p := newGoaResultView(res)
	return &rawimagesviews.GoaResult{Projected: p, View: "default"}
}

// newGoaResult converts projected type GoaResult to service type GoaResult.
func newGoaResult(vres *rawimagesviews.GoaResultView) *GoaResult {
	res := &GoaResult{
		Code:        vres.Code,
		Description: vres.Description,
	}
	return res
}

// newGoaResultView projects result type GoaResult to projected type
// GoaResultView using the "default" view.
func newGoaResultView(res *GoaResult) *rawimagesviews.GoaResultView {
	vres := &rawimagesviews.GoaResultView{
		Code:        res.Code,
		Description: res.Description,
	}
	return vres
}
