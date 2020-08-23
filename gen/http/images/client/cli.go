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
func BuildLoadNewSatelliteImagePayload(imagesLoadNewSatelliteImageBody string) (*images.SatelliteImage, error) {
	var err error
	var body LoadNewSatelliteImageRequestBody
	{
		err = json.Unmarshal([]byte(imagesLoadNewSatelliteImageBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"date_time\": \"1994-09-19T16:01:12Z\",\n      \"geographic_information\": {\n         \"coordinates\": {\n            \"Sit facilis qui et voluptas.\": 0.45731705303115716\n         },\n         \"tag_name\": \"Nemo occaecati aperiam quis autem.\"\n      },\n      \"id\": \"Harum architecto quo molestias qui.\",\n      \"normalized_indexes\": {\n         \"ndvi\": [\n            0.18711642837954914,\n            0.09177180070813358,\n            0.7808423638626252\n         ],\n         \"ndwi\": [\n            0.7146847364477503,\n            0.540422727294012,\n            0.5322622263700029,\n            0.299780450041929\n         ]\n      }\n   }'")
		}
	}
	v := &images.SatelliteImage{
		ID:       body.ID,
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
