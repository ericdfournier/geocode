package gmaps

import (
	"fmt"
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
)

// Test Connection Against Current IP
func GeocodeTestClientIP(clt *maps.Client) (e error) {

	var req maps.GeocodingRequest

	req = maps.GeocodingRequest{
		Address: "1600 Amphitheatre Pkwy, Mountain View, CA 94043",
	}

	_, err := clt.Geocode(context.Background(), &req)

	if err == nil {
		fmt.Println("Client IP Authenticated...")
	}

	return err
}
