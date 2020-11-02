// Code generated by goa v3.2.3, DO NOT EDIT.
//
// Processed images HTTP client encoders and decoders
//
// Command:
// $ goa gen image-loader/design

package client

import (
	"bytes"
	"context"
	processedimages "image-loader/gen/processed_images"
	processedimagesviews "image-loader/gen/processed_images/views"
	"io/ioutil"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
)

// BuildLoadNewProcessedSatelliteImageRequest instantiates a HTTP request
// object with method and path set to call the "Processed images" service "Load
// new processed satellite image" endpoint
func (c *Client) BuildLoadNewProcessedSatelliteImageRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: LoadNewProcessedSatelliteImageProcessedImagesPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("Processed images", "Load new processed satellite image", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeLoadNewProcessedSatelliteImageRequest returns an encoder for requests
// sent to the Processed images Load new processed satellite image server.
func EncodeLoadNewProcessedSatelliteImageRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*processedimages.ProcessedSatelliteImage)
		if !ok {
			return goahttp.ErrInvalidType("Processed images", "Load new processed satellite image", "*processedimages.ProcessedSatelliteImage", v)
		}
		body := NewLoadNewProcessedSatelliteImageRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("Processed images", "Load new processed satellite image", err)
		}
		return nil
	}
}

// DecodeLoadNewProcessedSatelliteImageResponse returns a decoder for responses
// returned by the Processed images Load new processed satellite image
// endpoint. restoreBody controls whether the response body should be restored
// after having been read.
// DecodeLoadNewProcessedSatelliteImageResponse may return the following errors:
//	- "BadRequest" (type *goa.ServiceError): http.StatusBadRequest
//	- "InternalError" (type *goa.ServiceError): http.StatusInternalServerError
//	- error: internal error
func DecodeLoadNewProcessedSatelliteImageResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body LoadNewProcessedSatelliteImageResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Processed images", "Load new processed satellite image", err)
			}
			p := NewLoadNewProcessedSatelliteImageGoaResultOK(&body)
			view := "default"
			vres := &processedimagesviews.GoaResult{Projected: p, View: view}
			if err = processedimagesviews.ValidateGoaResult(vres); err != nil {
				return nil, goahttp.ErrValidationError("Processed images", "Load new processed satellite image", err)
			}
			res := processedimages.NewGoaResult(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body LoadNewProcessedSatelliteImageBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Processed images", "Load new processed satellite image", err)
			}
			err = ValidateLoadNewProcessedSatelliteImageBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Processed images", "Load new processed satellite image", err)
			}
			return nil, NewLoadNewProcessedSatelliteImageBadRequest(&body)
		case http.StatusInternalServerError:
			var (
				body LoadNewProcessedSatelliteImageInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Processed images", "Load new processed satellite image", err)
			}
			err = ValidateLoadNewProcessedSatelliteImageInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Processed images", "Load new processed satellite image", err)
			}
			return nil, NewLoadNewProcessedSatelliteImageInternalError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("Processed images", "Load new processed satellite image", resp.StatusCode, string(body))
		}
	}
}

// BuildGetProcessedSatelliteImageRequest instantiates a HTTP request object
// with method and path set to call the "Processed images" service "Get
// processed satellite image" endpoint
func (c *Client) BuildGetProcessedSatelliteImageRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetProcessedSatelliteImageProcessedImagesPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("Processed images", "Get processed satellite image", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetProcessedSatelliteImageRequest returns an encoder for requests sent
// to the Processed images Get processed satellite image server.
func EncodeGetProcessedSatelliteImageRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*processedimages.GetProcessedSatelliteImagePayload)
		if !ok {
			return goahttp.ErrInvalidType("Processed images", "Get processed satellite image", "*processedimages.GetProcessedSatelliteImagePayload", v)
		}
		body := NewGetProcessedSatelliteImageRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("Processed images", "Get processed satellite image", err)
		}
		return nil
	}
}

// DecodeGetProcessedSatelliteImageResponse returns a decoder for responses
// returned by the Processed images Get processed satellite image endpoint.
// restoreBody controls whether the response body should be restored after
// having been read.
// DecodeGetProcessedSatelliteImageResponse may return the following errors:
//	- "BadRequest" (type *goa.ServiceError): http.StatusBadRequest
//	- "InternalError" (type *goa.ServiceError): http.StatusInternalServerError
//	- error: internal error
func DecodeGetProcessedSatelliteImageResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetProcessedSatelliteImageResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Processed images", "Get processed satellite image", err)
			}
			p := NewGetProcessedSatelliteImageGoaResultOK(&body)
			view := "default"
			vres := &processedimagesviews.GoaResult{Projected: p, View: view}
			if err = processedimagesviews.ValidateGoaResult(vres); err != nil {
				return nil, goahttp.ErrValidationError("Processed images", "Get processed satellite image", err)
			}
			res := processedimages.NewGoaResult(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body GetProcessedSatelliteImageBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Processed images", "Get processed satellite image", err)
			}
			err = ValidateGetProcessedSatelliteImageBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Processed images", "Get processed satellite image", err)
			}
			return nil, NewGetProcessedSatelliteImageBadRequest(&body)
		case http.StatusInternalServerError:
			var (
				body GetProcessedSatelliteImageInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Processed images", "Get processed satellite image", err)
			}
			err = ValidateGetProcessedSatelliteImageInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("Processed images", "Get processed satellite image", err)
			}
			return nil, NewGetProcessedSatelliteImageInternalError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("Processed images", "Get processed satellite image", resp.StatusCode, string(body))
		}
	}
}

// marshalProcessedimagesGeographicInformationToGeographicInformationRequestBody
// builds a value of type *GeographicInformationRequestBody from a value of
// type *processedimages.GeographicInformation.
func marshalProcessedimagesGeographicInformationToGeographicInformationRequestBody(v *processedimages.GeographicInformation) *GeographicInformationRequestBody {
	if v == nil {
		return nil
	}
	res := &GeographicInformationRequestBody{
		TagName: v.TagName,
	}
	if v.Coordinates != nil {
		res.Coordinates = make(map[string]float64, len(v.Coordinates))
		for key, val := range v.Coordinates {
			tk := key
			tv := val
			res.Coordinates[tk] = tv
		}
	}

	return res
}

// marshalProcessedimagesNormalizedIndexesToNormalizedIndexesRequestBody builds
// a value of type *NormalizedIndexesRequestBody from a value of type
// *processedimages.NormalizedIndexes.
func marshalProcessedimagesNormalizedIndexesToNormalizedIndexesRequestBody(v *processedimages.NormalizedIndexes) *NormalizedIndexesRequestBody {
	if v == nil {
		return nil
	}
	res := &NormalizedIndexesRequestBody{}
	if v.Ndvi != nil {
		res.Ndvi = make([]float64, len(v.Ndvi))
		for i, val := range v.Ndvi {
			res.Ndvi[i] = val
		}
	}
	if v.Ndwi != nil {
		res.Ndwi = make([]float64, len(v.Ndwi))
		for i, val := range v.Ndwi {
			res.Ndwi[i] = val
		}
	}

	return res
}

// marshalGeographicInformationRequestBodyToProcessedimagesGeographicInformation
// builds a value of type *processedimages.GeographicInformation from a value
// of type *GeographicInformationRequestBody.
func marshalGeographicInformationRequestBodyToProcessedimagesGeographicInformation(v *GeographicInformationRequestBody) *processedimages.GeographicInformation {
	if v == nil {
		return nil
	}
	res := &processedimages.GeographicInformation{
		TagName: v.TagName,
	}
	if v.Coordinates != nil {
		res.Coordinates = make(map[string]float64, len(v.Coordinates))
		for key, val := range v.Coordinates {
			tk := key
			tv := val
			res.Coordinates[tk] = tv
		}
	}

	return res
}

// marshalNormalizedIndexesRequestBodyToProcessedimagesNormalizedIndexes builds
// a value of type *processedimages.NormalizedIndexes from a value of type
// *NormalizedIndexesRequestBody.
func marshalNormalizedIndexesRequestBodyToProcessedimagesNormalizedIndexes(v *NormalizedIndexesRequestBody) *processedimages.NormalizedIndexes {
	if v == nil {
		return nil
	}
	res := &processedimages.NormalizedIndexes{}
	if v.Ndvi != nil {
		res.Ndvi = make([]float64, len(v.Ndvi))
		for i, val := range v.Ndvi {
			res.Ndvi[i] = val
		}
	}
	if v.Ndwi != nil {
		res.Ndwi = make([]float64, len(v.Ndwi))
		for i, val := range v.Ndwi {
			res.Ndwi[i] = val
		}
	}

	return res
}
