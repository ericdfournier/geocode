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
	"fmt"
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
	"gopkg.in/urfave/cli.v1"
)

// Establish Client API Connection
func ConnectClient(con *cli.Context) (clt *maps.Client, e error) {
    key := con.String("key")
    clt, err := maps.NewClient(maps.WithAPIKey(key))
    return clt, err
}

// Check API Connection Against Current IP
func CheckClientIP(con *cli.Context, clt *maps.Client) (e error) {
	// Allocations
	var (
		err     error  = nil
		success string = "Client IP Authenticated..."
		failure string = "Client IP Denied..."
	)
	// Switch on input command name
	switch con.Command.Name {
	case "geocode":
		// Allocate empty geocoder request object
		var req maps.GeocodingRequest
		// Build test request
		req = maps.GeocodingRequest{
			Address: "1600 Amphitheatre Pkwy, Mountain View, CA 94043",
		}
		// Submit test request
		_, err = clt.Geocode(context.Background(), &req)
	case "rvgeocode":
		// Allocate empty geocoder request object
		var req maps.GeocodingRequest
		// Build test request
		req = maps.GeocodingRequest{
			LatLng: &maps.LatLng{37.421915, -122.082699},
		}
		// Submit test request
		_, err = clt.Geocode(context.Background(), &req)
	case "elevation":
		// Allocate empty elevation request object
		var req maps.ElevationRequest
		// Build test request
		req = maps.ElevationRequest{
			Locations: []maps.LatLng{
				{
					Lat: 39.73915360,
					Lng: -104.9847034,
				},
			},
		}
		// Submit test request
		_, err = clt.Elevation(context.Background(), &req)
	case "nearby":
		// Allocate empty places request object
		var req maps.NearbySearchRequest
		// Build test request
		req = maps.NearbySearchRequest{
			Location: &maps.LatLng{39.73915360, -104.9847034},
			Radius:   1000,
		}
		// Submit test request
		_, err = clt.NearbySearch(context.Background(), &req)
	}
	// Print success message to console
	if err == nil {
		fmt.Println(success)
	} else {
		fmt.Println(failure)
	}
	return err
}
