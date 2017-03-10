package main

import(
        // "bytes"
	"fmt"
	"os"
        // "path/filepath"
)

func main() {
	
        arg := os.Args[1]
                
        junkDesc := "This is the Junk Command Line Tool /new Never 'rm -rf' your files or directories into oblivion ever again! /new Enter '--help' or '-h' for a list of options"
        
        if arg != "" {
                
                currDir, err := os.Getwd()
                homeDir, hasHomeDir := os.LookupEnv("HOME")
                junkDir := fmt.Sprint(homeDir + "/.Junk/")
                
                if err != nil {
                        fmt.Print("ERROR RETURNING CURRENT WORKING DIRECTORY")
                }
                
                goesToJunk := fmt.Sprint(currDir +"/"+ arg)
                newDir := fmt.Sprint(junkDir + arg)
                doesExist := Exists(goesToJunk)
                
                if doesExist && hasHomeDir {
                        os.Rename(goesToJunk, newDir)
                } else {
                        err := "File or directory does not exist in the current directory, ya dumb-butt."
                        fmt.Print(err)
                }
        } else {
                fmt.Print(junkDesc)
        }

}

//Checks whether a file or directory exists
func Exists(name string) bool {
    if _, err := os.Stat(name); err != nil {
    if os.IsNotExist(err) {
                return false
            }
    }
    return true
}
