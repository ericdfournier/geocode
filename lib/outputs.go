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
	"encoding/csv"
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"os"
	"strconv"
	"strings"
)

// Define Output Interface
type Output interface {
    Write() *csv.Writer
}

// Define fileOutput Struct
type fileOutput struct {
    path string
}

// Define consoleOutput Struct
type consoleOutput struct {
    stdin *os.File
}

// Define Write Method for fileOutput Struct
func (fp *fileOutput) Write() (file *os.File, writer *csv.Writer) {
        // Format output filepath
		out, err := OutputFilepath(fp.path)
		if err != nil {
			panic(err)
		}
		// Open output file
		f, err := os.Create(out)
		if err != nil {
			panic(err)
		}
		// Allocate new file writer
		w := csv.NewWriter(f)
        // Reader writer
        return f, w
}

// CSV Writer for Generating Output Elevation Results Files
func ElevationWriteOutput(con *cli.Context, results <-chan *ElevationRecord) (e error) {
	// Allocate empty error reciever
	var err error
	// Switch on output file flag
	switch con.IsSet("output") {
	case true:
        // Generate interface types
        fp := &fileOutput{con.String("output")}
        f, w := fp.Write()
        // Defer Closures
        defer f.Close()
        defer w.Flush()
        // Set writer formatting
		err = w.Write([]string{
			"id",
			"lat",
			"lng",
			"elevation",
			"resolution",
			"note"})
		if err != nil {
			panic(err)
		}
		// Enter Writer Loop
		lim := len(results)
		for i := 0; i < lim; i++ {
			// Extract current record from channel
			record := <-results
            //DEBUG 
            fmt.Println(record.Id)
			// Format strings
			latString := strconv.FormatFloat(record.Lat, 'f', -1, 64)
			lngString := strconv.FormatFloat(record.Lng, 'f', -1, 64)
			elevationString := strconv.FormatFloat(record.Elevation, 'f', -1, 64)
			resolutionString := strconv.FormatFloat(record.Resolution, 'f', -1, 64)
			// Write to file
			err = w.Write([]string{
				record.Id,
				latString,
				lngString,
				elevationString,
				resolutionString,
				record.Note})
			if err != nil {
				panic(err)
			}
		}
    default:
		// Enter writer loop
		lim := len(results)
		for i := 0; i < lim; i++ {
			// Extract current record from channel
			record := <-results
			// Format strings
			latString := strconv.FormatFloat(record.Lat, 'f', -1, 64)
			lngString := strconv.FormatFloat(record.Lng, 'f', -1, 64)
			elevationString := strconv.FormatFloat(record.Elevation, 'f', -1, 64)
			resolutionString := strconv.FormatFloat(record.Resolution, 'f', -1, 64)
			// Print to stdout
			values := []string{
				record.Id,
				latString,
				lngString,
				elevationString,
				resolutionString,
				record.Note}
			fmt.Println(strings.Join(values, ","))
		}
	}
	return err
}

// CSV Writer for Generating Geocoding Output Results Files
func GeocodeWriteOutput(con *cli.Context, results <-chan *GeocodeRecord) (e error) {
	// Allocated empty error reciever
	var err error = nil
	// Switch on output file flag
	switch con.IsSet("output") {
	case true:
        // Generate interface types
        fp := &fileOutput{con.String("output")}
        f, w := fp.Write()
		// Defer closures
		defer f.Close()
        defer w.Flush()
		// Format record outputs
		err = w.Write([]string{
			"id",
			"address",
			"lat",
			"lng",
			"note"})
		if err != nil {
			panic(err)
		}
		// Enter writer loop
		lim := len(results)
		for i := 0; i < lim; i++ {
			// Extract current record from channel
			record := <-results
			// Format strings
			latString := strconv.FormatFloat(record.Lat, 'f', -1, 64)
			lngString := strconv.FormatFloat(record.Lng, 'f', -1, 64)
			// Write to output file
			err := w.Write([]string{
				record.Id,
				record.Address,
				latString,
				lngString,
				record.Note})
			if err != nil {
				panic(err)
			}
		}
    default:
		// Enter writer loop
		lim := len(results)
		for i := 0; i < lim; i++ {
			// Extract current record from channel
			record := <-results
			// Format strings
			latString := strconv.FormatFloat(record.Lat, 'f', -1, 64)
			lngString := strconv.FormatFloat(record.Lng, 'f', -1, 64)
			// Print to stdout
			values := []string{
				record.Id,
				record.Address,
				latString,
				lngString,
				record.Note}
			fmt.Println(strings.Join(values, ","))
		}
	}
	return err
}

// CSV Writer for Generating Reverse Geocoding Output Results Files
func ReverseGeocodeWriteOutput(con *cli.Context, results <-chan *GeocodeRecord) (e error) {
	// Allocate empty error receiver
    var err error = nil
    // Switch on output file flag
    switch con.IsSet("output") {
    case true:
        // Generate interface types
        fp := &fileOutput{con.String("output")}
        f, w := fp.Write()
        // Defer closures
        defer f.Close()
        defer w.Flush()
        // Format record outputs
        err = w.Write([]string{
            "id",
            "lat",
            "lng",
            "address",
            "note"})
        if err != nil {
            panic(err)
        }
        // Enter writer loop
        lim := len(results)
        for i := 0; i < lim; i++ {
            // Extract current record from channel
            record := <-results
            // Format strings
            latString := strconv.FormatFloat(record.Lat, 'f', -1, 64)
            lngString := strconv.FormatFloat(record.Lng, 'f', -1, 64)
            // Write to output file
            err := w.Write([]string{
                record.Id,
                latString,
                lngString,
                record.Address,
                record.Note})
            if err != nil {
                panic(err)
            }
	    }
    default:
        // Enter writer loop
        lim := len(results)
        for i := 0; i < lim; i++ {
            // Extract current record from channel
            record := <-results
            // Format strings
            latString := strconv.FormatFloat(record.Lat, 'f', -1, 64)
            lngString := strconv.FormatFloat(record.Lng, 'f', -1, 64)
            // Write to output file
            values := []string{
                record.Id,
                latString,
                lngString,
                record.Address,
                record.Note}
            fmt.Println(strings.Join(values, ","))
	    }
    }
	return err
}

// CSV Writer for Generating Places Nearby Output Results Files
func PlaceNearbyWriteOutput(con *cli.Context, results <-chan *PlaceRecord) (e error) {
	// Allocate empty error receiver
    var err error = nil
    // Switch on output file flag
    switch con.IsSet("output") {
    case true:
        // Generate interface types
        fp := &fileOutput{con.String("output")}
        f, w := fp.Write()
        // Defer closure
        defer f.Close()
        defer w.Flush()
        // Format record outputs
        err = w.Write([]string{
            "id",
            "lat",
            "lng",
            "radius",
            "place_id",
            "name",
            "type",
            "note"})
        if err != nil {
            panic(err)
        }
        // Enter writer loop
        lim := len(results)
        for i := 0; i < lim; i++ {
            // Extract current record from channel
            record := <-results
            // Format strings
            latString := strconv.FormatFloat(record.Lat, 'f', -1, 64)
            lngString := strconv.FormatFloat(record.Lng, 'f', -1, 64)
            radiusString := strconv.Itoa(int(record.Radius))
            // Write to output file
            err := w.Write([]string{
                record.Id,
                latString,
                lngString,
                radiusString,
                record.PlaceId,
                record.Name,
                record.Type,
                record.Note})
            if err != nil {
                panic(err)
            }
        }
    default:
        // Enter writer loop
        lim := len(results)
        for i := 0; i < lim; i++ {
            // Extract current record from channel
            record := <-results
            // Format strings
            latString := strconv.FormatFloat(record.Lat, 'f', -1, 64)
            lngString := strconv.FormatFloat(record.Lng, 'f', -1, 64)
            radiusString := strconv.Itoa(int(record.Radius))
            // Write to output file
            values := []string{
                record.Id,
                latString,
                lngString,
                radiusString,
                record.PlaceId,
                record.Name,
                record.Type,
                record.Note}
            fmt.Println(strings.Join(values, ","))
        }
    }
	return err
}
