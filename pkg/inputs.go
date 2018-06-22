package gmaps

import (
	"bufio"
	"encoding/csv"
	"os"
	"strconv"
)

// CSV Reader for Processing Elevation Input Files
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
				Id:  record[0],
				Lat: latFloat,
				Lng: lngFloat,
			}
		}
	}
	return records, err
}

// CSV Reader for Processing Geocoding Input Files
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

// CSV Reader for Processing Place Nearby Input Files
func PlaceNearbyReadCSV(filepath string) (output chan *PlaceNearbyRecord, e error) {
	// Open input file
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	// Defer closure
	defer f.Close()
	r := csv.NewReader(bufio.NewReader(f))
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
			// Parse radius to int
			radiusInt, err := strconv.Atoi(record[3])
			if err != nil {
				panic(err)
			}
			records <- &PlaceNearbyRecord{Id: record[0], Lat: latFloat, Lng: lngFloat, Radius: uint(radiusInt)}
		}
	}
	return records, err
}

// CSV Reader for Processing Place Details Input Files
func PlaceDetailsReadCSV(filepath string) (output chan *PlaceDetailRecord, e error) {
	// Open input file
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	// Defer closure
	defer f.Close()
	r := csv.NewReader(bufio.NewReader(f))
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
		if i == 0 {
			continue
		} else {
			records <- &PlaceDetailRecord{Id: record[0], PlaceId: record[1]}
		}
	}
	return records, err
}
