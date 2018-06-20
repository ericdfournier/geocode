package gmaps

// Geocode Record Struct Field Specification
type GeocodeRecord struct {
	Id      string
	Address string
	Lat     float64
	Lng     float64
	Region  string
	Note    string
}

// Elevation Record Struct Field Spedification
type ElevationRecord struct {
	Id         string
	Elevation  float64
	Lat        float64
	Lng        float64
	Resolution float64
	Note       string
}
