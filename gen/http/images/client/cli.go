// Code generated by goa v3.2.3, DO NOT EDIT.
//
// Images HTTP client CLI support package
//
// Command:
// $ goa gen image-loader/design

package client

import (
	"encoding/json"
	"fmt"
	images "image-loader/gen/images"
)

// BuildLoadNewSatelliteImagePayload builds the payload for the Images Load new
// satellite image endpoint from CLI flags.
func BuildLoadNewSatelliteImagePayload(imagesLoadNewSatelliteImageBody string) (*images.RawSatelliteImage, error) {
	var err error
	var body LoadNewSatelliteImageRequestBody
	{
		err = json.Unmarshal([]byte(imagesLoadNewSatelliteImageBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"file_name\": \"Voluptates architecto.\",\n      \"id\": \"Consequuntur non placeat ab.\"\n   }'")
		}
	}
	v := &images.RawSatelliteImage{
		ID:       body.ID,
		FileName: body.FileName,
	}

	return v, nil
}

// BuildGetRawSatelliteImagePayload builds the payload for the Images Get raw
// satellite image endpoint from CLI flags.
func BuildGetRawSatelliteImagePayload(imagesGetRawSatelliteImageBody string) (*images.GetRawSatelliteImagePayload, error) {
	var err error
	var body GetRawSatelliteImageRequestBody
	{
		err = json.Unmarshal([]byte(imagesGetRawSatelliteImageBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"file_name\": \"Quia quos eos ut unde sed dolorum.\"\n   }'")
		}
	}
	v := &images.GetRawSatelliteImagePayload{
		FileName: body.FileName,
	}

	return v, nil
}

// BuildLoadNewProcessedSatelliteImagePayload builds the payload for the Images
// Load new processed satellite image endpoint from CLI flags.
func BuildLoadNewProcessedSatelliteImagePayload(imagesLoadNewProcessedSatelliteImageBody string) (*images.ProcessedSatelliteImage, error) {
	var err error
	var body LoadNewProcessedSatelliteImageRequestBody
	{
		err = json.Unmarshal([]byte(imagesLoadNewProcessedSatelliteImageBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"date_time\": \"2008-09-18T10:35:05Z\",\n      \"file_name\": \"Reprehenderit tempore nesciunt aut.\",\n      \"geographic_information\": {\n         \"coordinates\": {\n            \"Ex exercitationem aut reiciendis quod voluptate incidunt.\": 0.7392343776263494,\n            \"Molestiae sed molestias praesentium rerum.\": 0.3908551358201754,\n            \"Qui qui iusto praesentium.\": 0.19181675265671314\n         },\n         \"tag_name\": \"Incidunt odit tempore.\"\n      },\n      \"id\": \"Sunt occaecati ut qui libero similique dolores.\",\n      \"normalized_indexes\": {\n         \"ndvi\": [\n            0.6127802172492277,\n            0.004560762359582058,\n            0.480367361557016\n         ],\n         \"ndwi\": [\n            0.08253636129242008,\n            0.030755381445945546,\n            0.659276994468877\n         ]\n      }\n   }'")
		}
	}
	v := &images.ProcessedSatelliteImage{
		ID:       body.ID,
		FileName: body.FileName,
		DateTime: body.DateTime,
	}
	if body.GeographicInformation != nil {
		v.GeographicInformation = marshalGeographicInformationRequestBodyToImagesGeographicInformation(body.GeographicInformation)
	}
	if body.NormalizedIndexes != nil {
		v.NormalizedIndexes = marshalNormalizedIndexesRequestBodyToImagesNormalizedIndexes(body.NormalizedIndexes)
	}

	return v, nil
}
