package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
	"gopkg.in/cheggaaa/pb.v1"
	"gopkg.in/urfave/cli.v1"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"time"
)

// Record Struct Field Specification
type Record struct {
	Id      string
	Address string
	Lat     float64
	Lng     float64
    Region  string
	Note    string
}

// CSV Reader for Processing Input Files
func csvReader(filepath string) (output chan *Record, e error) {
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	r := csv.NewReader(bufio.NewReader(f))
	r.Comma = ','
	r.FieldsPerRecord = -1

	rawData, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	records := make(chan *Record, len(rawData))

	for i, record := range rawData {

		if i == 0 {
			continue
		} else {
			records <- &Record{Id: record[0], Address: record[1]}
		}
	}

	return records, nil
}

// CSV Writer for Generating Output Results Files
func csvWriter(filepath string, results <-chan *Record) (e error) {

	f, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	err = w.Write([]string{"id", "address", "lat", "lng", "note"})
	if err != nil {
		panic(err)
	}

	lim := len(results)

	for i := 0; i < lim; i++ {

		record := <-results

		latString := strconv.FormatFloat(record.Lat, 'f', -1, 64)
		lngString := strconv.FormatFloat(record.Lng, 'f', -1, 64)
		err := w.Write([]string{record.Id, record.Address, latString, lngString, record.Note})
		if err != nil {
			panic(err)
		}

	}

	return nil
}

// Test Connection Against Current IP
func testClientIP(clt *maps.Client) (e error) {

    var req maps.GeocodingRequest

    req = maps.GeocodingRequest{
        Address: "1600 Amphitheatre Pkwy, Mountain View, CA 94043",
    }

    _, err := clt.Geocode(context.Background(), &req)

    return err
}

// Record Formatting Method for Assembling the API Request
func formatRequest(con *cli.Context, rec *Record) (request maps.GeocodingRequest) {

    var req maps.GeocodingRequest

    if len(con.String("region")) == 0 {
        req = maps.GeocodingRequest{
            Address: rec.Address,
        }
    } else {
        req = maps.GeocodingRequest{
            Address: rec.Address,
            Region: con.String("region"),
        }
    }

    return req
}

// Wrapper Function to Automate the API Calls
func geocodeRecords(con *cli.Context, clt *maps.Client, records <-chan *Record) (results chan *Record, e error) {

	results = make(chan *Record, len(records))
	lim := len(records)
	bar := pb.StartNew(lim)

	for i := 0; i < lim; i++ {

		rec := <-records
        req := formatRequest(con, rec)

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

	return results, nil
}

// Function for Parsing Command Line Arguments
func checkArgs(con *cli.Context) (e error) {

	if len(con.String("input")) == 0 {
		return cli.NewExitError("ERROR: Must Provide Input Filepath", 1)
	}

	_, err := os.Stat(con.String("input"))

	if os.IsNotExist(err) {
		return cli.NewExitError("ERROR: Input Filepath Does Not Exist", 2)
	}

	if len(con.String("key")) == 0 {
		return cli.NewExitError("ERROR: Must Provide Valid API Key", 3)
	}

	return nil
}

// Function for Formatting API Call Response
func formatOutput(con *cli.Context) (out string, e error) {

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

// Global Variables for CLI
var apiKey string = ""
var input string = ""
var output string = ""
var region string = ""

// Main Function
func main() {

	geocode := cli.NewApp()
	geocode.Name = "geocode"
	geocode.Version = "00.01.2"
	geocode.Compiled = time.Now()
	geocode.Authors = []cli.Author{
		cli.Author{
			Name:  "Eric Daniel Fournier",
			Email: "me@ericdfournier.com",
		},
	}
	geocode.HelpName = "geocode"
	geocode.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "key, k",
			Usage: "Google Maps API 'Key'",
			Value: apiKey,
		},
		cli.StringFlag{
			Name:  "input, i",
			Usage: "Input 'Filepath'",
			Value: input,
		},
		cli.StringFlag{
			Name:  "output, o",
			Usage: "Output 'Directory Path'",
			Value: output,
		},
        cli.StringFlag{
            Name: "region, r",
            Usage: "Restricted 'Region Code'",
            Value: region,
        },
	}

	sort.Sort(cli.FlagsByName(geocode.Flags))

	geocode.Action = func(con *cli.Context) (e error) {

		err := checkArgs(con)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

        clt, err := maps.NewClient(maps.WithAPIKey(con.String("key")))
		if err != nil {
			fmt.Println(err)
            os.Exit(2)
		}

        err = testClientIP(clt)
        if err != nil {
            fmt.Println(err)
            os.Exit(2)
        }

		out, err := formatOutput(con)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		rec, err := csvReader(con.String("input"))
        if err != nil {
            fmt.Println(err)
            os.Exit(2)
        }

		res, err  := geocodeRecords(con, clt, rec)
        if err != nil {
            fmt.Println(err)
            os.Exit(2)
        }

        err = csvWriter(out, res)
        if err != nil{
            fmt.Println(err)
            os.Exit(2)
        }

		return err
	}

	geocode.Run(os.Args)
	os.Exit(1)
}
