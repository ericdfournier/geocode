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