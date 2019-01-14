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

// Wrapper function to Automate Reverse Geocoding API Calls
func ReverseGeocodeRecords(con *cli.Context, clt *maps.Client, records <-chan *GeocodeRecord) (results chan *GeocodeRecord, e error) {
	// Allocate empty variables
	var err error = nil
    // Allocate receiver variables
	lim := len(records)
    results = make(chan *GeocodeRecord, lim)
	bar := pb.StartNew(lim)
	// Enter request loop
	for i := 0; i < lim; i++ {
		//Extract current record
		rec := <-records
		req := ReverseGeocodeFormatRequest(con, rec)
		// Submit requests and process errors
		if req.LatLng.Lat != 0 && req.LatLng.Lng != 0 {
			res, err := clt.Geocode(context.Background(), &req)
			if err != nil {
				fmt.Println(err)
			}
			if len(res) != 0 {
				rec.Address = res[0].FormattedAddress
				rec.Note = "Success"
			} else {
				rec.Note = "No Reverse Geocoding Result"
			}
		} else {
			rec.Note = "Lat and/or Lng Missing"
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
    // Allocate receiver variables
	lim := len(records)
    results = make(chan *ElevationRecord, lim)
	bar := pb.StartNew(lim)
	// Enter request loop
	for i := 0; i < lim; i++ {
		// Extract current records
		rec := <-records
		req := ElevationFormatRequest(con, rec)
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
func PlaceNearbyRecords(con *cli.Context, clt *maps.Client, records <-chan *PlaceRecord) (results chan *PlaceRecord, e error) {
	// Allocate empty variables
	var err error = nil
    // Allocate receiver variables
	lim := len(records)
    results = make(chan *PlaceRecord, lim)
	bar := pb.StartNew(lim)
	// Enter request loop
	for i := 0; i < lim; i++ {
		// Extract current records
		rec := <-records
		req := PlaceNearbyFormatRequest(con, rec)
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
				if len(res.Results) > 1 {
					rec.Note = "Success: Multiple Place Results Found - First Retrieved"
				} else {
					rec.Note = "Success"
				}
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

func PlaceDetailRecords(con *cli.Context, clt *maps.Client, records <-chan *PlaceRecord) (results chan *PlaceRecord, e error){
    // Allocate empty variables
    var err error = nil
    // Allocate reciever variables
    lim := len(records)
    results = make(chan *PlaceRecord, lim)
    bar := pb.StartNew(lim)
    // Enter request loop
    for i := 0; i < lim; i++ {
        rec := <-records
        req := PlaceDetailFormatRequest(con, rec)
        // Submit requests and process errors
        if req.PlaceID != "" {
            res, err := clt.PlaceDetails(context.Background(), &req)
            if err != nil {
                fmt.Println(err)
            }
                rec.Name = res.Name
                rec.Scope = res.Scope
                rec.Type = res.Types[0]
                rec.Viewport = res.Geometry.Viewport
                rec.Bounds = res.Geometry.Bounds
        } else {
            rec.Note = "Place ID Missing"
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
