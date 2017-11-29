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

type Record struct {
	Id      string
	Address string
	Lat     float64
	Lng     float64
	Note    string
}

func csvReader(filepath string) (output chan *Record) {
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

	return records
}

func csvWriter(filepath string, results <-chan *Record) {

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

	return
}

func geocodeRecords(c *maps.Client, records <-chan *Record) (results chan *Record) {

	results = make(chan *Record, len(records))
	lim := len(records)
	bar := pb.StartNew(lim)

	for i := 0; i < lim; i++ {

		record := <-records

		req := &maps.GeocodingRequest{
			Address: record.Address,
		}

		if req.Address != "" {
			res, err := c.Geocode(context.Background(), req)
			if err != nil {
				fmt.Println(err)
			}
			if len(res) != 0 {
				record.Lat = res[0].Geometry.Location.Lat
				record.Lng = res[0].Geometry.Location.Lng
				record.Note = "Success"
			} else {
				record.Note = "No Geocoding Result"
			}
		} else {
			record.Note = "Address Missing"
		}
		results <- record
		bar.Increment()
	}

	bar.Finish()

	return results
}

var apiKey string = ""

func main() {

	geocode := cli.NewApp()
	geocode.Name = "geocode"
	geocode.Version = "00.01.0"
	geocode.Compiled = time.Now()
	geocode.Authors = []cli.Author{
		cli.Author{
			Name:  "Eric Daniel Fournier",
			Email: "me@ericdfournier.com",
		},
	}
	geocode.HelpName = "geocode"

	t := time.Now().Format(time.RFC3339)

	p1, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	input := filepath.Join(p1, "records.csv")

	p2, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	output := filepath.Join(p2, "results_"+t+".csv")

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
			Usage: "Output 'Filepath'",
			Value: output,
		},
	}

	sort.Sort(cli.FlagsByName(geocode.Flags))

	geocode.Action = func(c *cli.Context) error {

		if len(c.String("key")) == 0 {
			return cli.NewExitError("ERROR: Must Provide Valid API Key", 69)
		}

		records := csvReader(c.String("input"))

		client, err := maps.NewClient(maps.WithAPIKey(c.String("key")))
		if err != nil {
			panic(err)
		}

		results := geocodeRecords(client, records)

		csvWriter(c.String("output"), results)

		return err
	}

	geocode.Run(os.Args)
	os.Exit(1)
}
