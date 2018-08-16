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

// Define Input Interface
type Input interface {
	Read() *csv.Reader
}

// Define Filepath Input Struct
type Filepath struct {
	path string
}

// Define Console Input Struct
type Console struct {
	stdin *os.File
}

// Define Read Method for Filepath Input
func (fp *Filepath) Read() *csv.Reader {
	// Open input file
	f, err := os.Open(fp.path)
	if err != nil {
		panic(err)
	}
	// Defer closure
	defer f.Close()
	// Allocate new file reader
	r := csv.NewReader(bufio.NewReader(f))
	// Parameterize reader
	r.Comma = ','
	r.FieldsPerRecord = -1
	// Return reader
	return r
}

// Define Read Method for Console Input
func (cs *Console) Read() *csv.Reader {
	// Read from stdin
	r := csv.NewReader(cs.stdin)
	// Return reader
	return r
}

// Reader for Processing Elevation Inputs
func ElevationReadInput(con *cli.Context) (output chan *ElevationRecord, e error) {
	// Allocate empty reader
	var r *csv.Reader = nil
	// Switch on context input
	switch con.IsSet("input") {
	case true:
		fp := &Filepath{con.String("input")}
		r = fp.Read()
	default:
		cs := &Console{os.Stdin}
		r = cs.Read()
	}
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
	// Allocate empty reader
	var r *csv.Reader = nil
	// Switch on context input
	switch con.IsSet("input") {
	case true:
		fp := &Filepath{con.String("input")}
		r = fp.Read()
	default:
		cs := &Console{os.Stdin}
		r = cs.Read()
	}
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
	// Switch on context input
	switch con.IsSet("input") {
	case true:
		fp := &Filepath{con.String("input")}
		r = fp.Read()
	default:
		cs := &Console{os.Stdin}
		r = cs.Read()
	}
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
func PlaceNearbyReadInput(con *cli.Context) (output chan *PlaceRecord, e error) {
	// Allocate empty reader
	var r *csv.Reader = nil
	// Switch on context input
	switch con.IsSet("input") {
	case true:
		fp := &Filepath{con.String("input")}
		r = fp.Read()
	default:
		cs := &Console{os.Stdin}
		r = cs.Read()
	}
	// Read input records
	rawData, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	// Allocate empty records channel
	records := make(chan *PlaceRecord, len(rawData))
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
			records <- &PlaceRecord{
				Id:     record[0],
				Lat:    latFloat,
				Lng:    lngFloat,
				Radius: uint(radiusInt)}
		}
	}
	return records, err
}

// Reader for Processing Place Detail Inputs
func PlaceDetailsReadInput(con *cli.Context) (output chan *PlaceRecord, e error) {
	// Allocate empty reader
	var r *csv.Reader = nil
	// Switch on context input
	switch con.IsSet("input") {
	case true:
		fp := &Filepath{con.String("input")}
		r = fp.Read()
	default:
		cs := &Console{os.Stdin}
		r = cs.Read()
	}
	// Read input records
	rawData, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	// Allocate empty records channel
	records := make(chan *PlaceRecord, len(rawData))
	// Enter record channel population loop
	for i, record := range rawData {
		// Skip header row
		if i == 0 && con.IsSet("input") {
			continue
		} else {
			records <- &PlaceRecord{
				Id:      record[0],
				PlaceId: record[1]}
		}
	}
	return records, err
}
