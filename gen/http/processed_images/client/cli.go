// Code generated by goa v3.2.3, DO NOT EDIT.
//
// Processed images HTTP client CLI support package
//
// Command:
// $ goa gen image-loader/design

package client

import (
	"encoding/json"
	"fmt"
	processedimages "image-loader/gen/processed_images"
)

// BuildLoadNewProcessedSatelliteImagePayload builds the payload for the
// Processed images Load new processed satellite image endpoint from CLI flags.
func BuildLoadNewProcessedSatelliteImagePayload(processedImagesLoadNewProcessedSatelliteImageBody string) (*processedimages.ProcessedSatelliteImage, error) {
	var err error
	var body LoadNewProcessedSatelliteImageRequestBody
	{
		err = json.Unmarshal([]byte(processedImagesLoadNewProcessedSatelliteImageBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"date_time\": \"2012-04-04T19:21:12Z\",\n      \"file_name\": \"Sunt aliquam totam dolores aut voluptatem sunt.\",\n      \"geographic_information\": {\n         \"coordinates\": {\n            \"Dolores aut reprehenderit tempore nesciunt.\": 0.08537234062709087,\n            \"Incidunt odit tempore.\": 0.4612094935047014\n         },\n         \"tag_name\": \"Ut qui.\"\n      },\n      \"normalized_indexes\": {\n         \"ndvi\": [\n            0.3904972148805288,\n            0.6699739812364446,\n            0.027588443100172985\n         ],\n         \"ndwi\": [\n            0.508680136375543,\n            0.9307709997993568\n         ]\n      }\n   }'")
		}
	}
	v := &processedimages.ProcessedSatelliteImage{
		FileName: body.FileName,
		DateTime: body.DateTime,
	}
	if body.GeographicInformation != nil {
		v.GeographicInformation = marshalGeographicInformationRequestBodyToProcessedimagesGeographicInformation(body.GeographicInformation)
	}
	if body.NormalizedIndexes != nil {
		v.NormalizedIndexes = marshalNormalizedIndexesRequestBodyToProcessedimagesNormalizedIndexes(body.NormalizedIndexes)
	}

	return v, nil
}

// BuildGetProcessedSatelliteImagePayload builds the payload for the Processed
// images Get processed satellite image endpoint from CLI flags.
func BuildGetProcessedSatelliteImagePayload(processedImagesGetProcessedSatelliteImageBody string) (*processedimages.GetProcessedSatelliteImagePayload, error) {
	var err error
	var body GetProcessedSatelliteImageRequestBody
	{
		err = json.Unmarshal([]byte(processedImagesGetProcessedSatelliteImageBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"file_name\": \"Voluptatem nihil.\"\n   }'")
		}
	}
	v := &processedimages.GetProcessedSatelliteImagePayload{
		FileName: body.FileName,
	}

	return v, nil
}
