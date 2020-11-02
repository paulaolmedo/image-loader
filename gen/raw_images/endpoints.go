// Code generated by goa v3.2.3, DO NOT EDIT.
//
// Raw images endpoints
//
// Command:
// $ goa gen image-loader/design

package rawimages

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "Raw images" service endpoints.
type Endpoints struct {
	LoadNewRawSatelliteImage goa.Endpoint
	GetRawSatelliteImage     goa.Endpoint
}

// NewEndpoints wraps the methods of the "Raw images" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		LoadNewRawSatelliteImage: NewLoadNewRawSatelliteImageEndpoint(s),
		GetRawSatelliteImage:     NewGetRawSatelliteImageEndpoint(s),
	}
}

// Use applies the given middleware to all the "Raw images" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.LoadNewRawSatelliteImage = m(e.LoadNewRawSatelliteImage)
	e.GetRawSatelliteImage = m(e.GetRawSatelliteImage)
}

// NewLoadNewRawSatelliteImageEndpoint returns an endpoint function that calls
// the method "Load new raw satellite image" of service "Raw images".
func NewLoadNewRawSatelliteImageEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*RawSatelliteImage)
		res, err := s.LoadNewRawSatelliteImage(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedGoaResult(res, "default")
		return vres, nil
	}
}

// NewGetRawSatelliteImageEndpoint returns an endpoint function that calls the
// method "Get raw satellite image" of service "Raw images".
func NewGetRawSatelliteImageEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetRawSatelliteImagePayload)
		res, err := s.GetRawSatelliteImage(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedGoaResult(res, "default")
		return vres, nil
	}
}
