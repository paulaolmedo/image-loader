// Code generated by goa v3.2.3, DO NOT EDIT.
//
// Images client
//
// Command:
// $ goa gen image-loader/design

package images

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "Images" service client.
type Client struct {
	LoadNewSatelliteImageEndpoint goa.Endpoint
}

// NewClient initializes a "Images" service client given the endpoints.
func NewClient(loadNewSatelliteImage goa.Endpoint) *Client {
	return &Client{
		LoadNewSatelliteImageEndpoint: loadNewSatelliteImage,
	}
}

// LoadNewSatelliteImage calls the "Load new satellite image" endpoint of the
// "Images" service.
// LoadNewSatelliteImage may return the following errors:
//	- "ErrorAddingImage" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) LoadNewSatelliteImage(ctx context.Context, p *SatelliteImage) (err error) {
	_, err = c.LoadNewSatelliteImageEndpoint(ctx, p)
	return
}