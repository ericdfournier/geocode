package gmaps

import (
	"fmt"
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
	"gopkg.in/cheggaaa/pb.v1"
	"gopkg.in/urfave/cli.v1"
)

// Wrapper Function to Automate the API Calls
func GeocodeRecords(con *cli.Context, clt *maps.Client, records <-chan *GeocodeRecord) (results chan *GeocodeRecord, e error) {
	// Allocate empty variables
	var err error = nil
	results = make(chan *GeocodeRecord, len(records))
	lim := len(records)
	bar := pb.StartNew(lim)
	// Enter request loop
	for i := 0; i < lim; i++ {
		// Extract current record
		rec := <-records
		req := GeocodeFormatRequest(con, rec)
		// Submit requests and process errors
		if req.Address != "" {
			res, err := clt.Geocode(context.Background(), &req)
			if err != nil {
				fmt.Println(err)
			}
			if len(res) != 0 {
				rec.Lat = res[0].Geometry.Location.Lat
				rec.Lng = res[0].Geometry.Location.Lng
				rec.Note = "Success"
			} else {
				rec.Note = "No Geocoding Result"
			}
		} else {
			rec.Note = "Address Missing"
		}
		// Send results to channel
		results <- rec
		// Increment progress bar
		bar.Increment()
	}
	// Finish progress bar
	bar.Finish()
	return results, err
}

// Wrapper Function to Automate Elevation API Calls
func ElevationRecords(con *cli.Context, clt *maps.Client, records <-chan *ElevationRecord) (results chan *ElevationRecord, e error) {
	// Allocate empty variables
	var err error = nil
	results = make(chan *ElevationRecord, len(records))
	lim := len(records)
	bar := pb.StartNew(lim)
	// Enter request loop
	for i := 0; i < lim; i++ {
		// Extract current records
		rec := <-records
		req := ElevationFormatRequest(rec)
		// Submit requests and process errors
		if req.Locations[0].Lat != 0 && req.Locations[0].Lng != 0 {
			res, err := clt.Elevation(context.Background(), &req)
			if err != nil {
				fmt.Println(err)
			}
			if len(res) != 0 {
				rec.Elevation = res[0].Elevation
				rec.Resolution = res[0].Resolution
				rec.Note = "Success"
			} else {
				rec.Note = "No Elevation Result"
			}
		} else {
			rec.Note = "Latitude or Longitude Missing"
		}
		// Send results to channel
		results <- rec
		// Increment progress bar
		bar.Increment()
	}
	// Finish progress bar
	bar.Finish()
	return results, err
}

// Wrapper Function to Automate Places API Nearby Calls
func PlaceNearbyRecords(con *cli.Context, clt *maps.Client, records <-chan *PlaceNearbyRecord) (results chan *PlaceNearbyRecord, e error) {
	// Allocate empty variables
	var err error = nil
	results = make(chan *PlaceNearbyRecord, len(records))
	lim := len(records)
	bar := pb.StartNew(lim)
	// Enter request loop
	for i := 0; i < lim; i++ {
		// Extract current records
		rec := <-records
		req := PlaceNearbyFormatRequest(rec)
		// Submit requests and process errors
		if req.Location.Lat != 0 && req.Location.Lng != 0 {
			res, err := clt.NearbySearch(context.Background(), &req)
			if err != nil {
				fmt.Println(err)
			}
			if len(res.Results) != 0 {
				rec.PlaceId = res.Results[0].PlaceID
				rec.Name = res.Results[0].Name
				rec.Type = res.Results[0].Types[0]
				rec.Note = "Success"
			} else {
				rec.Note = "No Place Result"
			}
		} else {
			rec.Note = "Latitude or Longitude Missing"
		}
		// Send results to channel
		results <- rec
		// Increment progress bar
		bar.Increment()
	}
	// Finish progress bar
	bar.Finish()
	return results, err
}
