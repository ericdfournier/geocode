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
	"bufio"
	"encoding/csv"
	"gopkg.in/urfave/cli.v1"
	"os"
	"strconv"
)

// Reader for Processing Elevation Inputs
func ElevationReadInput(con *cli.Context) (output chan *ElevationRecord, e error) {
	// Allocate empty reader
	var r *csv.Reader = nil
	// Switch on input flag
	if con.IsSet("input") {
		// Open input file
		f, err := os.Open(con.String("input"))
		if err != nil {
			panic(err)
		}
		// Defer closure
		defer f.Close()
		// Allocate new file reader
		r = csv.NewReader(bufio.NewReader(f))
	} else {
		// Read from stdin
		r = csv.NewReader(os.Stdin)
	}
	// Parameterize reader
	r.Comma = ','
	r.FieldsPerRecord = -1
	// Read in the raw data
	rawData, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	// Allocate empty records channel
	records := make(chan *ElevationRecord, len(rawData))
	// Enter reader loops
	for i, record := range rawData {
		// Skip header row
		if i == 0 && con.IsSet("input") {
			continue
		} else {
			// Parse lat float
			latFloat, err := strconv.ParseFloat(record[1], 64)
			if err != nil {
				panic(err)
			}
			// Parse lon float
			lngFloat, err := strconv.ParseFloat(record[2], 64)
			if err != nil {
				panic(err)
			}
			// Send formatted record to channel
			records <- &ElevationRecord{
				Id:  record[0],
				Lat: latFloat,
				Lng: lngFloat,
			}
		}
	}
	return records, err
}

// Reader for Processing Geocoding Inputs
func GeocodeReadInput(con *cli.Context) (output chan *GeocodeRecord, e error) {
	// Allocate reader reciever
	var r *csv.Reader = nil
	// Switch on input type
	if con.IsSet("input") {
		// Open input file
		f, err := os.Open(con.String("input"))
		if err != nil {
			panic(err)
		}
		// Defer closure
		defer f.Close()
		// Allocated new csv reader
		r = csv.NewReader(bufio.NewReader(f))
	} else {
		// Read from stdin
		r = csv.NewReader(os.Stdin)
	}
	// Parameterize reader
	r.Comma = ','
	r.FieldsPerRecord = -1
	// Read input records
	rawData, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	// Allocate empty records channel
	records := make(chan *GeocodeRecord, len(rawData))
	// Enter record channel population loop
	for i, record := range rawData {
		// Skip header row
		if i == 0 && con.IsSet("input") == true {
			continue
		} else {
			records <- &GeocodeRecord{
				Id:      record[0],
				Address: record[1]}
		}
	}
	return records, err
}

// Reader for Processing Reverse Geocoding Inputs
func ReverseGeocodeReadInput(con *cli.Context) (output chan *GeocodeRecord, e error) {
	// Allocate empty reader
	var r *csv.Reader = nil
	// Switch on input type
	if con.IsSet("input") {
		// Open input file
		f, err := os.Open(con.String("input"))
		if err != nil {
			panic(err)
		}
		// Defer closure
		defer f.Close()
		// Allocated new csv reader
		r = csv.NewReader(bufio.NewReader(f))
	} else {
		// Read from stdin
		r = csv.NewReader(os.Stdin)
	}
	// Parameterize reader
	r.Comma = ','
	r.FieldsPerRecord = -1
	// Read input records
	rawData, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	// Allocate empty records channel
	records := make(chan *GeocodeRecord, len(rawData))
	// Enter record channel population loop
	for i, record := range rawData {
		// Skip header row
		if i == 0 && con.IsSet("input") {
			continue
		} else {
			// Parse lat float
			latFloat, err := strconv.ParseFloat(record[1], 64)
			if err != nil {
				panic(err)
			}
			// Parse lng float
			lngFloat, err := strconv.ParseFloat(record[2], 64)
			if err != nil {
				panic(err)
			}
			// Write to record
			records <- &GeocodeRecord{
				Id:  record[0],
				Lat: latFloat,
				Lng: lngFloat}
		}
	}
	return records, err
}

// Reader for Processing Place Nearby Inputs
func PlaceNearbyReadInput(con *cli.Context) (output chan *PlaceNearbyRecord, e error) {
	// Allocate empty reader
	var r *csv.Reader = nil
	// Switch on input type
	if con.IsSet("input") {
		// Open input file
		f, err := os.Open(con.String("input"))
		if err != nil {
			panic(err)
		}
		// Defer closure
		defer f.Close()
		r = csv.NewReader(bufio.NewReader(f))
	} else {
		// Read from stdin
		r = csv.NewReader(os.Stdin)
	}
	// Parameterize reader
	r.Comma = ','
	r.FieldsPerRecord = -1
	// Read input records
	rawData, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	// Allocate empty records channel
	records := make(chan *PlaceNearbyRecord, len(rawData))
	// Enter record channel population loop
	for i, record := range rawData {
		// Skip header row
		if i == 0 && con.IsSet("input") {
			continue
		} else {
			// Parse lat float
			latFloat, err := strconv.ParseFloat(record[1], 64)
			if err != nil {
				panic(err)
			}
			// Parse lon float
			lngFloat, err := strconv.ParseFloat(record[2], 64)
			if err != nil {
				panic(err)
			}
			// Parse radius to int
			radiusInt, err := strconv.Atoi(record[3])
			if err != nil {
				panic(err)
			}
			records <- &PlaceNearbyRecord{
				Id:     record[0],
				Lat:    latFloat,
				Lng:    lngFloat,
				Radius: uint(radiusInt)}
		}
	}
	return records, err
}

// Reader for Processing Place Detail Inputs
func PlaceDetailsReadInput(con *cli.Context) (output chan *PlaceDetailRecord, e error) {
	// Allocated empty csv reader
	var r *csv.Reader = nil
	// Switch on input type
	if con.IsSet("input") {
		// Open input file
		f, err := os.Open(con.String("input"))
		if err != nil {
			panic(err)
		}
		// Defer closure
		defer f.Close()
		r = csv.NewReader(bufio.NewReader(f))
	} else {
		r = csv.NewReader(os.Stdin)
	}
	// Parameterize csv reader
	r.Comma = ','
	r.FieldsPerRecord = -1
	// Read input records
	rawData, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	// Allocate empty records channel
	records := make(chan *PlaceDetailRecord, len(rawData))
	// Enter record channel population loop
	for i, record := range rawData {
		// Skip header row
		if i == 0 && con.IsSet("input") {
			continue
		} else {
			records <- &PlaceDetailRecord{
				Id:      record[0],
				PlaceId: record[1]}
		}
	}
	return records, err
}
