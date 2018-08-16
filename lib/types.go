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
