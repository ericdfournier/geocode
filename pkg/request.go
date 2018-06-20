package gmaps

import (
	"fmt"
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
	"gopkg.in/cheggaaa/pb.v1"
	"gopkg.in/urfave/cli.v1"
)

// Wrapper Function to Automate the API Calls
func GeocodeRecords(con *cli.Context, clt *maps.Client, records <-chan *GeoRecord) (results chan *GeoRecord, e error) {

	var err error = nil
	results = make(chan *GeoRecord, len(records))
	lim := len(records)
	bar := pb.StartNew(lim)

	for i := 0; i < lim; i++ {

		rec := <-records
		req := GeocodeFormatRequest(con, rec)

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

		results <- rec
		bar.Increment()

	}

	bar.Finish()

	return results, err
}
