// Code generated by goa v3.2.3, DO NOT EDIT.
//
// Raw images HTTP server encoders and decoders
//
// Command:
// $ goa gen image-loader/design

package server

import (
	"context"
	rawimagesviews "image-loader/gen/raw_images/views"
	"io"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeLoadNewRawSatelliteImageResponse returns an encoder for responses
// returned by the Raw images Load new raw satellite image endpoint.
func EncodeLoadNewRawSatelliteImageResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*rawimagesviews.GoaResult)
		enc := encoder(ctx, w)
		body := NewLoadNewRawSatelliteImageResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeLoadNewRawSatelliteImageRequest returns a decoder for requests sent to
// the Raw images Load new raw satellite image endpoint.
func DecodeLoadNewRawSatelliteImageRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body LoadNewRawSatelliteImageRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		payload := NewLoadNewRawSatelliteImageRawSatelliteImage(&body)

		return payload, nil
	}
}

// EncodeLoadNewRawSatelliteImageError returns an encoder for errors returned
// by the Load new raw satellite image Raw images endpoint.
func EncodeLoadNewRawSatelliteImageError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "BadRequest":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewLoadNewRawSatelliteImageBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", "BadRequest")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "InternalError":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewLoadNewRawSatelliteImageInternalErrorResponseBody(res)
			}
			w.Header().Set("goa-error", "InternalError")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeGetRawSatelliteImageResponse returns an encoder for responses returned
// by the Raw images Get raw satellite image endpoint.
func EncodeGetRawSatelliteImageResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*rawimagesviews.GoaResult)
		enc := encoder(ctx, w)
		body := NewGetRawSatelliteImageResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetRawSatelliteImageRequest returns a decoder for requests sent to the
// Raw images Get raw satellite image endpoint.
func DecodeGetRawSatelliteImageRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body GetRawSatelliteImageRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		payload := NewGetRawSatelliteImagePayload(&body)

		return payload, nil
	}
}

// EncodeGetRawSatelliteImageError returns an encoder for errors returned by
// the Get raw satellite image Raw images endpoint.
func EncodeGetRawSatelliteImageError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "BadRequest":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewGetRawSatelliteImageBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", "BadRequest")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "InternalError":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewGetRawSatelliteImageInternalErrorResponseBody(res)
			}
			w.Header().Set("goa-error", "InternalError")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}
