// Code generated by goa v3.2.3, DO NOT EDIT.
//
// Images endpoints
//
// Command:
// $ goa gen image-loader/design

package images

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "Images" service endpoints.
type Endpoints struct {
	LoadNewSatelliteImage goa.Endpoint
}

// NewEndpoints wraps the methods of the "Images" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		LoadNewSatelliteImage: NewLoadNewSatelliteImageEndpoint(s),
	}
}

// Use applies the given middleware to all the "Images" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.LoadNewSatelliteImage = m(e.LoadNewSatelliteImage)
}

// NewLoadNewSatelliteImageEndpoint returns an endpoint function that calls the
// method "Load new satellite image" of service "Images".
func NewLoadNewSatelliteImageEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SatelliteImage)
		return nil, s.LoadNewSatelliteImage(ctx, p)
	}
}
