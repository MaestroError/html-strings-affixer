package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/MaestroError/html-strings-affixer/app"
	"github.com/MaestroError/html-strings-affixer/reporter"
)

// Entry point
func main() {
	app.Bootstrap()
	
	createConfigFileIfRequested()

	app.Start()
}


// Checks config file existence and creates it if requested from user
func createConfigFileIfRequested() {
	// Set main vars
	configFileName := "affixer.json"
	reporter := reporter.Reporter{}
	// Get command run directly (pwd)
	path := pwd() + "/" + configFileName

	// check if config file exists
	configExists, err := exists(path)
	if err != nil {
		panic(err)
	}

	if !configExists {
		// get Executable location
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		exPath := filepath.Dir(ex)
		// get config example path
		exampleFile := exPath + "/affixer-example.json"
		exampleExists, err := exists(exampleFile)
		if err != nil {
			panic(err)
		}

		// check if example config file exists
		if exampleExists {
			// Ask user to create
			if reporter.AskForConfirmation("You have no affixer.json config file. Do you wanna create it? (Will not affect this run)", "yes") {
				// Read example file and get content
				var r []byte
				var err error
				r, err = ioutil.ReadFile(exampleFile)
				if err != nil {
					panic(err)
				}
				content := string(r)
				// write content in new configFileName file in pwd
				errWrite := ioutil.WriteFile(path, []byte(content), 0)
				if errWrite != nil {
					panic(err)
				}
			}
		}
	}

}

// Checks if file exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// Get current working directory (where run the command)
func pwd() string {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return path
}