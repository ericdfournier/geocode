package gmaps

import (
	"bufio"
	"encoding/csv"
	"os"
	"strconv"
)

// CSV Reader for Processing Input Files
func GeocodeReadCSV(filepath string) (output chan *GeoRecord, e error) {
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

	records := make(chan *GeoRecord, len(rawData))

	for i, record := range rawData {

		if i == 0 {
			continue
		} else {
			records <- &GeoRecord{Id: record[0], Address: record[1]}
		}
	}

	return records, err
}

// CSV Writer for Generating Output Results Files
func GeocodeWriteCSV(filepath string, results <-chan *GeoRecord) (e error) {

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

	return err
}
