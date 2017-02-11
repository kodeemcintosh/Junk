package main

import(
	"os"
	"fmt"
        "github.com/kvmac/Junk/flags.go"
)

func main() {
	
	moveJunk()

}
flags.initializer(args, junkDesc)
args := os.Args
cliArg := os.Args[0]
currDir := os.Getwd()
homeDir := os.LookupEnv("HOME")
junkDesc := "This is the Junk Command Line Tool"/new"Never 'rm -rf' your files or directories into oblivion ever again!"/new"Enter '--help' or '-h' for a list of options"
junkDir := "/.Junk"
oldPath := currDir+cliArg


//func switchEnv(toThisDir) {

//	os.Chdir(toThisDir+junkDir)
//	currDir = toThisDir
	

//}

func moveJunk() {

	os.Rename(oldPath,newPath)

}


//THIS GOES INTO A SEPARATE FILE NAMED flags.go
////////|||||||
////////vvvvvvv

args := os.Args
argsLen := len(args)
initializer(args, argsLen)

func initializer(argsArray []string, num int) {
        if (num > 4) {
                Println("Error: There should be no more than four arguments")
        }
        else if (num == 4) {
                argFour := os.Args[3]
        }
        else if(num == 3) {
                argThree := os.Args[2]
        }
        else if (num == 2) {
                argTwo := os.Args[1]
        }
        else if (num == 1) {
                argOne := os.Args[0]
        }
        else {
                appOverview()
        }
}

func appOverview() {

Println("This is the Junk command line tool")
Println("Never rm -rf another directory into oblivion, again!")
Println("Use the argument '--help' or '-h' to obtain a list of options")

}

func filterArg() (string) {


}
