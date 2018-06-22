package gmaps

import (
	"gopkg.in/urfave/cli.v1"
	"os/user"
	"path/filepath"
)

// Function for Formatting API Call Response Output Filepath
func OutputFilepath(con *cli.Context) (out string, e error) {
	// Allocate vars
	var err error = nil
	var output string
	// Parse command line arguments
	if len(con.String("output")) == 0 {
		// Get user info
		usr, err := user.Current()
		if err != nil {
			output = "None"
			return output, err
		}
		// Set default output to user home directory
		output = filepath.Join(usr.HomeDir, "results.csv")
	} else if len(con.String("output")) != 0 {
		// Format output filename
		output = con.String("output")
	}
	return output, err
}
