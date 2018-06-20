package gmaps

import (
	"googlemaps.github.io/maps"
	"gopkg.in/urfave/cli.v1"
)

// Format Geocode Record for API Request
func GeocodeFormatRequest(con *cli.Context, rec *GeocodeRecord) (request maps.GeocodingRequest) {
	// Allocated empty request
	var req maps.GeocodingRequest
	// Set request properties on optional flags
	if len(con.String("region")) == 0 {
		req = maps.GeocodingRequest{
			Address: rec.Address,
		}
	} else {
		req = maps.GeocodingRequest{
			Address: rec.Address,
			Region:  con.String("region"),
		}
	}
	return req
}

// Format Elevation Record for API Request
func ElevationFormatRequest(rec *ElevationRecord) (request maps.ElevationRequest) {
	// Allocated empty request
	var req maps.ElevationRequest
	// Set request format
	req = maps.ElevationRequest{
		Locations: []maps.LatLng{
			{
				Lat: rec.Lat,
				Lng: rec.Lng,
			},
		},
	}
	return req
}
