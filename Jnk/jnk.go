package main

import (
	"fmt"
	"os"
)

func main() {
	a := os.Args

	if len(a) == 2 {
		moveThings(a[2])

	} else if len(a) > 2 {
		iterate(a)

	} else {
		fmt.Println(junkDesc)

	}

}

// Iterates over multiple command line arguments
func iterate(args []string) {
	for _, a := range args {
		moveThings(a)

	}

}

func moveThings(arg string) {
	// Variables for directory paths
	currDir, err := os.Getwd()
	if err != nil {
		fmt.Print("ERROR RETURNING CURRENT WORKING DIRECTORY")
	}
	homeDir, hasHomeDir := os.LookupEnv("HOME")
	junkDir := fmt.Sprint(homeDir + "/.Junk/")
	newDir := fmt.Sprint(junkDir + arg)
	goesToJunk := fmt.Sprint(currDir + "/" + arg)

	// Bool variable to make sure that the file/directory path exists
	doesExist := exists(goesToJunk)

	if doesExist && hasHomeDir {
		os.Rename(goesToJunk, newDir)

	} else {
		err := "File or directory does not exist in the current directory, ya dumb-butt."
		fmt.Print(err)

	}

}

var junkDesc string = "This is the Junk Command Line Tool! Never 'rm -rf' your files or directories into oblivion ever again! Enter '--help' or '-h' for a list of options"

//Checks whether a file or directory exists
func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
