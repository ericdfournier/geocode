package gmaps

import (
	"googlemaps.github.io/maps"
	"gopkg.in/urfave/cli.v1"
	"os"
	"path/filepath"
	"time"
)

// Record Formatting Method for Assembling the API Request
func GeocodeFormatRequest(con *cli.Context, rec *GeoRecord) (request maps.GeocodingRequest) {

	var req maps.GeocodingRequest

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

// Function for Formatting API Call Response
func GeocodeFormatResponse(con *cli.Context) (out string, e error) {

	t := time.Now().Format(time.RFC3339)
	var err error = nil
	var output string

	if len(con.String("output")) == 0 {

		p, err := os.Getwd()
		if err != nil {
			output = "None"
			return output, err
		}

		output = filepath.Join(p, "results_"+t+".csv")

	} else if len(con.String("output")) != 0 {

		_, err := os.Stat(con.String("output"))

		if os.IsNotExist(err) {
			output = "None"
			err = cli.NewExitError("ERROR: Output Directory Does Not Exist", 5)
			return output, err
		}

		output = filepath.Join(con.String("output"), "results_"+t+".csv")

	}

	return output, err
}
