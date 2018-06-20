package gmaps

import (
	"fmt"
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
)

// Test Geocoder API Connection Against Current IP
func GeocodeTestClientIP(clt *maps.Client) (e error) {
	// Allocate empty geocoder request object
	var req maps.GeocodingRequest
	// Build test request
	req = maps.GeocodingRequest{
		Address: "1600 Amphitheatre Pkwy, Mountain View, CA 94043",
	}
	// Submit test request
	_, err := clt.Geocode(context.Background(), &req)
	// Print success message to console
	if err == nil {
		fmt.Println("Client IP Authenticated Against User Geocoder API Key...")
	}
	return err
}

// Test Elevation API Connection Against Current IP
func ElevationTestClientIP(clt *maps.Client) (e error) {
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
	_, err := clt.Elevation(context.Background(), &req)
	// Print success message to console
	if err == nil {
		fmt.Println("Client IP Authenticated Against User Elevation API Key...")
	}
	return err
}
