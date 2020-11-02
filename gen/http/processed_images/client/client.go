// Code generated by goa v3.2.3, DO NOT EDIT.
//
// Processed images client HTTP transport
//
// Command:
// $ goa gen image-loader/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the Processed images service endpoint HTTP clients.
type Client struct {
	// LoadNewProcessedSatelliteImage Doer is the HTTP client used to make requests
	// to the Load new processed satellite image endpoint.
	LoadNewProcessedSatelliteImageDoer goahttp.Doer

	// GetProcessedSatelliteImage Doer is the HTTP client used to make requests to
	// the Get processed satellite image endpoint.
	GetProcessedSatelliteImageDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the Processed images service
// servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		LoadNewProcessedSatelliteImageDoer: doer,
		GetProcessedSatelliteImageDoer:     doer,
		RestoreResponseBody:                restoreBody,
		scheme:                             scheme,
		host:                               host,
		decoder:                            dec,
		encoder:                            enc,
	}
}

// LoadNewProcessedSatelliteImage returns an endpoint that makes HTTP requests
// to the Processed images service Load new processed satellite image server.
func (c *Client) LoadNewProcessedSatelliteImage() goa.Endpoint {
	var (
		encodeRequest  = EncodeLoadNewProcessedSatelliteImageRequest(c.encoder)
		decodeResponse = DecodeLoadNewProcessedSatelliteImageResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildLoadNewProcessedSatelliteImageRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.LoadNewProcessedSatelliteImageDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("Processed images", "Load new processed satellite image", err)
		}
		return decodeResponse(resp)
	}
}

// GetProcessedSatelliteImage returns an endpoint that makes HTTP requests to
// the Processed images service Get processed satellite image server.
func (c *Client) GetProcessedSatelliteImage() goa.Endpoint {
	var (
		encodeRequest  = EncodeGetProcessedSatelliteImageRequest(c.encoder)
		decodeResponse = DecodeGetProcessedSatelliteImageResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetProcessedSatelliteImageRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetProcessedSatelliteImageDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("Processed images", "Get processed satellite image", err)
		}
		return decodeResponse(resp)
	}
}
