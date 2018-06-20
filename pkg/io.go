package gmaps

import (
	"bufio"
	"encoding/csv"
    "gopkg.in/urfave/cli.v1"
	"os"
	"strconv"
    "time"
    "path/filepath"
)

// Function for Formatting API Call Response Output Filepath
func OutputFilepath(con *cli.Context) (out string, e error) {
    // Get current time for label
	t := time.Now().Format(time.RFC3339)
    // Allocate vars
	var err error = nil
	var output string
    // Parse command line arguments
	if len(con.String("output")) == 0 {
		p, err := os.Getwd()
		if err != nil {
			output = "None"
			return output, err
		}
        // Format output filename
		output = filepath.Join(p, "results_"+t+".csv")
	} else if len(con.String("output")) != 0 {
		_, err := os.Stat(con.String("output"))
		if os.IsNotExist(err) {
			output = "None"
			err = cli.NewExitError("ERROR: Output Directory Does Not Exist", 5)
			return output, err
		}
        // Format output filename
		output = filepath.Join(con.String("output"), "results_"+t+".csv")

	}
	return output, err
}

// CSV Reader for Processing Input Files
func ElevationReadCSV(filepath string) (output chan *ElevationRecord, e error) {
	// Open input file
    f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
    // Defer closure
	defer f.Close()
    // Allocate new file reader
	r := csv.NewReader(bufio.NewReader(f))
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
		if i == 0 {
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
                Id: record[0],
                Lat: latFloat,
                Lng: lngFloat,
            }
		}
	}
	return records, err
}

// CSV Writer for Generating Output Elevation Results Files
func ElevationWriteCSV(filepath string, results <-chan *ElevationRecord) (e error) {
    // Open output file
	f, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
    // Defer closure
	defer f.Close()
    // Allocate new file writer
	w := csv.NewWriter(f)
	defer w.Flush()
    // Set writer formatting
	err = w.Write([]string{"id", "lat", "lng", "elevation", "resolution", "note"})
	if err != nil {
		panic(err)
	}
    // Enter Writer Loop
	lim := len(results)
	for i := 0; i < lim; i++ {
        // Extract current record from channel
		record := <-results
        // Format strings
		latString := strconv.FormatFloat(record.Lat, 'f', -1, 64)
		lngString := strconv.FormatFloat(record.Lng, 'f', -1, 64)
        elevationString := strconv.FormatFloat(record.Elevation, 'f', -1, 64)
        resolutionString := strconv.FormatFloat(record.Resolution, 'f', -1, 64)
		// Write to file
        err := w.Write([]string{record.Id, latString, lngString, elevationString, resolutionString, record.Note})
		if err != nil {
			panic(err)
		}
	}
	return err
}

// CSV Reader for Processing Input Files
func GeocodeReadCSV(filepath string) (output chan *GeocodeRecord, e error) {
	// Open input file
    f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
    // Defer closure
	defer f.Close()
    // Allocated new csv reader
	r := csv.NewReader(bufio.NewReader(f))
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
		if i == 0 {
			continue
		} else {
			records <- &GeocodeRecord{Id: record[0], Address: record[1]}
		}
	}
	return records, err
}

// CSV Writer for Generating Output Results Files
func GeocodeWriteCSV(filepath string, results <-chan *GeocodeRecord) (e error) {
    // Open output file
	f, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
    // Defer closure
	defer f.Close()
    // Allocate empty writer
	w := csv.NewWriter(f)
	defer w.Flush()
    // Format record outputs
	err = w.Write([]string{"id", "address", "lat", "lng", "note"})
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
        err := w.Write([]string{record.Id, record.Address, latString, lngString, record.Note})
		if err != nil {
			panic(err)
		}
	}
	return err
}
