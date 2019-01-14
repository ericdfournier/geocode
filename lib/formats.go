/*
Copyright (c) 2018 Eric Daniel Fournier

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

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

// Format Reverse Geocode Record for API Request
func ReverseGeocodeFormatRequest(con *cli.Context, rec *GeocodeRecord) (request maps.GeocodingRequest) {
	// Allocate empty request
	var req maps.GeocodingRequest
	//Set request properties on optional flags
	if len(con.String("region")) == 0 {
		req = maps.GeocodingRequest{
			LatLng: &maps.LatLng{rec.Lat, rec.Lng},
		}
	} else {
		req = maps.GeocodingRequest{
			LatLng: &maps.LatLng{rec.Lat, rec.Lng},
			Region: con.String("region"),
		}
	}
	return req
}

// Format Elevation Record for API Request
func ElevationFormatRequest(con *cli.Context, rec *ElevationRecord) (request maps.ElevationRequest) {
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

// Format Place Nearby Record for API Request
func PlaceNearbyFormatRequest(con *cli.Context, rec *PlaceRecord) (request maps.NearbySearchRequest) {
	// Allocated empty request
	var req maps.NearbySearchRequest
	// Set request format
	req = maps.NearbySearchRequest{
		Location: &maps.LatLng{rec.Lat, rec.Lng},
		Radius:   rec.Radius,
	}
	return req
}

// Format Place Detail Record for API Request
func PlaceDetailFormatRequest(con *cli.Context, rec *PlaceRecord) (request maps.PlaceDetailsRequest) {
    // Allocated empty request
    var req maps.PlaceDetailsRequest
    // Set request format
    req = maps.PlaceDetailsRequest{
        PlaceID: rec.Id,
    }
    return req
}
