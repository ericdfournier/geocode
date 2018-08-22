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

package main

import (
	"fmt"
	gm "github.com/ericdfournier/gmaps/lib"
	"googlemaps.github.io/maps"
	"gopkg.in/urfave/cli.v1"
	"os"
	"sort"
	"time"
)

// Global Variables for CLI
var apiKey string = ""
var input string = ""
var output string = ""
var region string = ""

// Function to Check if Inputs Can Be Sourced from Piped Stdin
func CheckCharDevice() (info os.FileInfo) {
	// Get stdin stat
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	// If piped inputs not supported print error message
	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: echo 'input' | gmaps [...]")
		panic(err)
	}
	return info
}

// Function for Parsing Command Line Arguments
func CheckArgs(con *cli.Context) (e error) {
	// Allocated empty error
	var err error
	// Get stdin
	info := CheckCharDevice()
	// Check if input flag is set
	if con.IsSet("input") != true && info.Size() <= 0 {
		return cli.NewExitError("ERROR: Must Recieve STDIN or Provide Input Filepath", 1)
	}
	// Check if input file exists
	if con.IsSet("input") == true {
		_, err := os.Stat(con.String("input"))
		if os.IsNotExist(err) {
			return cli.NewExitError("ERROR: Input Filepath Does Not Exist", 2)
		}
	}
	// Check if api key flag is set
	if con.IsSet("key") != true {
		return cli.NewExitError("ERROR: Must Provide Valid API Key", 3)
	}
	return err
}

// Main Function
func main() {
	gmaps := cli.NewApp()
	gmaps.Name = "gmaps"
	gmaps.Usage = "Command Line Interface to Google Maps Web Service APIs"
	gmaps.Version = "00.06.03"
	gmaps.Compiled = time.Now()
	gmaps.Authors = []cli.Author{
		cli.Author{
			Name:  "Eric Daniel Fournier",
			Email: "me@ericdfournier.com",
		},
	}
	gmaps.HelpName = "gmaps"
	gmaps.Commands = []cli.Command{
		// Geocoder API Sub-Command
		{
			Name:  "geocode",
			Usage: "Google Maps Geocoder API Tool",
			Description: `
			Accepts STDIN or Input FILEPATH [CSV].
			Outputs STDOUT or Output FILEPATH [CSV].
			Input STDIN Format: 
				id - [string],
				address - [string]
			Output STDOUT Format:
				... ,
				lat - [float],
				lng - [float],
				note - [string]`,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "key, k",
					Usage:  "Google Maps Geocoder API 'Key'",
					Value:  apiKey,
					EnvVar: "GMAPS_API_KEY",
				},
				cli.StringFlag{
					Name: "input, i",
					Usage: `
					Input FILEPATH Format: 
						id - [string], 
						address - [string]`,
					Value: input,
				},
				cli.StringFlag{
					Name: "output, o",
					Usage: `
					Output FILEPATH Format: 
						... ,
						lat - [float],
						lng - [float],
						note - [string]`,
					Value: output,
				},
				cli.StringFlag{
					Name:  "region, r",
					Usage: "Restricted 'Region Code'",
					Value: region,
				},
			},
			Action: func(con *cli.Context) (e error) {
				// Check input arguments
				err := CheckArgs(con)
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
				// Establish new Google Maps API client connection
				clt, err := maps.NewClient(maps.WithAPIKey(con.String("key")))
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
				// Authenticate client IP
				err = gm.CheckClientIP(con, clt)
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
				// Read in address data from csv file
				rec, err := gm.GeocodeReadInput(con)
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
				// Geocode records from input csv file records
				res, err := gm.GeocodeRecords(con, clt, rec)
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
				// Write formatted output to csv file
				err = gm.GeocodeWriteOutput(con, res)
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
				return err
			},
		},
		// Reverse Geocoder API Sub-Command
		{
			Name:  "rvgeocode",
			Usage: "Google Maps Reverse Geocoder API Tool",
			Description: `
			Accepts STDIN or Input FILEPATH [CSV]. 
			Outputs STDOUT or Output FILEPATH [CSV].
			Input STDIN Format: 
				id - [string],
				lat - [float],
				lng - [float]
			Output STDOUT Format:
				... ,
				address - [string],
				note - [string]`,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "key, k",
					Usage:  "Google Maps Reverse Geocoder API 'Key'",
					Value:  apiKey,
					EnvVar: "GMAPS_API_KEY",
				},
				cli.StringFlag{
					Name: "input, i",
					Usage: `
					Input FILEPATH Format: 
						id - [string], 
						lat - [float], 
						lng - [float]`,
					Value: input,
				},
				cli.StringFlag{
					Name: "output, o",
					Usage: `
					Output FILEPATH Format: 
						... ,
						address - [string], 
						note - [string]`,
					Value: output,
				},
				cli.StringFlag{
					Name:  "region, r",
					Usage: "Restricted 'Region Code'",
					Value: region,
				},
			},
			Action: func(con *cli.Context) (e error) {
				// Check input arguments
				err := CheckArgs(con)
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
				// Establish new Google Maps API client connection
				clt, err := maps.NewClient(maps.WithAPIKey(con.String("key")))
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
				// Authenticate client IP
				err = gm.CheckClientIP(con, clt)
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
				// Read in address data from csv file
				rec, err := gm.ReverseGeocodeReadInput(con)
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
				// Geocode records from input csv file records
				res, err := gm.ReverseGeocodeRecords(con, clt, rec)
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
				// Write formatted output to csv file
				err = gm.ReverseGeocodeWriteOutput(con, res)
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
				return err
			},
		},
		{
			Name:  "place",
			Usage: "Google Maps Places API Tool",
			Description: `Options for searching for nearby places or accessing 
			detailed place information from the Google Maps Place API.`,
			Subcommands: []cli.Command{
				{
					Name:  "nearby",
					Usage: "Search for nearby places by latitude, longitude",
					Description: `
					Accepts STDIN or Input FILEPATH [CSV].
					Outputs STDOUT or Output FILEPATH [CSV].
					Input STDIN Format: 
						id - [string],
						lat - [float],
						lng - [float],
					Output STDOUT Format:
						...	,
						placeId - [string],
						name - [string],
						type - [string],
						note - [string]`,
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:   "key, k",
							Usage:  "Google Place API 'Key'",
							Value:  apiKey,
							EnvVar: "GMAPS_API_KEY",
						},
						cli.StringFlag{
							Name: "input, i",
							Usage: `
							Input FILEPATH Format:
								id - [string],
								lat - [float],
								lng - [float],
								radius - [int]`,
							Value: input,
						},
						cli.StringFlag{
							Name: "output, o",
							Usage: `
							Output FILEPATH Format:
								... ,
								placeId - [string],
								name - [string],
								type - [string],
								note - [string]`,
							Value: output,
						},
					},
					Action: func(con *cli.Context) (e error) {
						// Check input arguments
						err := CheckArgs(con)
						if err != nil {
							fmt.Println(err)
							os.Exit(2)
						}
						// Establish new Google Maps API client connections
						clt, err := maps.NewClient(maps.WithAPIKey(con.String("key")))
						if err != nil {
							fmt.Println(err)
							os.Exit(2)
						}
						// Authenticate client IP
						err = gm.CheckClientIP(con, clt)
						if err != nil {
							fmt.Println(err)
							os.Exit(2)
						}
						// Read in coordinate data from csv file
						rec, err := gm.PlaceNearbyReadInput(con)
						if err != nil {
							fmt.Println(err)
							os.Exit(2)
						}
						// Request place data from input CSV file records
						res, err := gm.PlaceNearbyRecords(con, clt, rec)
						if err != nil {
							fmt.Println(err)
							os.Exit(2)
						}
						// Write formatted output to CSV file
						err = gm.PlaceNearbyWriteOutput(con, res)
						if err != nil {
							fmt.Println(err)
							os.Exit(2)
						}
						return err
					},
				},
				{
					Name:  "detail",
					Usage: "Search for specific details by google place ID",
					Description: `Accests an Input CSV File With Formated Google
					Location IDs and Outputs a Formatted CSV File with Placed
					Response Details.`,
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:   "key, k",
							Usage:  "Google Maps Places API 'Key'",
							Value:  apiKey,
							EnvVar: "GMAPS_API_KEY",
						},
						cli.StringFlag{
							Name: "input, i",
							Usage: `
							Input Format:
								id - [string],
								placeId - [string]`,
							Value: input,
						},
						cli.StringFlag{
							Name: "output, o",
							Usage: `
							Output Format:
								...,
								TBD`,
							Value: output,
						},
					},
					Action: func(con *cli.Context) (e error) {
						// Check input arguments
						err := CheckArgs(con)
						if err != nil {
							fmt.Println(err)
							os.Exit(2)
						}
						return
					},
				},
			},
		},
		{
			Name:  "elevation",
			Usage: "Google Maps Elevation API Tool",
			Description: `
			Accepts STDIN or Input FILEPATH [CSV]. 
			Outputs STDOUT or Output FILEPATH [CSV].
			Input STDIN Format: 
				id - [string],
				lat - [float],
				lng - [float]
			Output STDOUT Format:
				... ,
				elevation - [float],
				resolution - [float],
				note - [string]`,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "key, k",
					Usage:  "Google Maps Elevation API 'Key'",
					Value:  apiKey,
					EnvVar: "GMAPS_API_KEY",
				},
				cli.StringFlag{
					Name: "input, i",
					Usage: `
					Input FILEPATH Format:
						id - [string],
						lat - [float],
						lng - [float]`,
					Value: input,
				},
				cli.StringFlag{
					Name: "output, o",
					Usage: `
					Output FILEPATH Format:
						... - ,
						elevation - [float],
						resolution - [float],
						note - [string]`,
					Value: output,
				},
			},
			Action: func(con *cli.Context) (e error) {
				// Check input arguments
				err := CheckArgs(con)
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
				// Establish new Google Maps API client connection
				clt, err := maps.NewClient(maps.WithAPIKey(con.String("key")))
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
				// Authenticate client IP
				err = gm.CheckClientIP(con, clt)
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
				// Read in coordinate data from csv file
				rec, err := gm.ElevationReadInput(con)
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
				// Request elevations from input csv file records
				res, err := gm.ElevationRecords(con, clt, rec)
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
				// Write formatted output to csv file
				err = gm.ElevationWriteOutput(con, res)
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
				return err
			},
		},
	}
	sort.Sort(cli.FlagsByName(gmaps.Flags))
	gmaps.Run(os.Args)
	os.Exit(1)
}
