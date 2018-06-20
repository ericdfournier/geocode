package main

import (
	"fmt"
	gm "github.com/ericdfournier/gmaps/pkg"
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

// Function for Parsing Command Line Arguments
func CheckArgs(con *cli.Context) (e error) {
	if con.IsSet("input") != true {
		return cli.NewExitError("ERROR: Must Provide Input Filepath", 1)
	}
	_, err := os.Stat(con.String("input"))
	if os.IsNotExist(err) {
		return cli.NewExitError("ERROR: Input Filepath Does Not Exist", 2)
	}
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
	gmaps.Version = "00.03.0"
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
            Name: "geocode",
			Aliases: []string{"gc"},
			Usage: "Google Maps Geocoder API Tool",
			Description: `Accepts an Input CSV File With Address Strings and 
			Outputs Formatted CSV File With Geocoder Response 
			Latitude, Longitude Results.`,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "key, k",
					Usage: "Google Maps Geocoder API 'Key'",
					Value: apiKey,
                    EnvVar: "GMAPS_API_KEY",
				},
				cli.StringFlag{
					Name: "input, i",
					Usage: "Input 'Filepath'",
					Value: input,
				},
				cli.StringFlag{
					Name: "output, o",
					Usage: "Output 'Directory Path'",
					Value: output,
				},
				cli.StringFlag{
					Name: "region, r",
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
				err = gm.GeocodeTestClientIP(clt)
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
				// Read in address data from csv file
				rec, err := gm.GeocodeReadCSV(con.String("input"))
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
                // Format output filepath
				out, err := gm.OutputFilepath(con)
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
				// Write formatted output to csv file
				err = gm.GeocodeWriteCSV(out, res)
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
				return err
			},
		},
        {
            Name: "elevation",
            Aliases: []string{"ev"},
            Usage: "Google Maps Elevation API Tool",
            Description: `Accepts an Input CSV File With Latitude, Longitude
            Pairs and Outputs a Formatted CSV File with Elevation Response
            Results`,
            Flags: []cli.Flag{
                cli.StringFlag{
                    Name: "key, k",
                    Usage: "Google Maps Elevation API 'Key'",
                    Value: apiKey,
                    EnvVar: "GMAPS_API_KEY",
                },
                cli.StringFlag{
                    Name: "input, i",
                    Usage: "Input 'Filepath'",
                    Value: output,
                },
                cli.StringFlag{
                    Name: "output, o",
                    Usage: "Output 'Directory Path'",
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
                if err != nil{
                    fmt.Println(err)
                    os.Exit(2)
                }
                // Authenticate client IP
                err = gm.ElevationTestClientIP(clt)
                if err != nil {
                    fmt.Println(err)
                    os.Exit(2)
                }
                // Read in coordinate data from csv file
                rec, err := gm.ElevationReadCSV(con.String("input"))
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
                // Format output filepath
				out, err := gm.OutputFilepath(con)
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
                // Write formatted output to csv file
                err = gm.ElevationWriteCSV(out, res)
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
