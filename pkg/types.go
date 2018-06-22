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

// Place Nearby Record Struct Field Specification
type PlaceNearbyRecord struct {
	Id      string
	Lat     float64
	Lng     float64
	Radius  uint
	PlaceId string
	Name    string
	Type    string
	Note    string
}

// Place Details Record Struct Field Specification
type PlaceDetailRecord struct {
	Id      string
	PlaceId string
}
