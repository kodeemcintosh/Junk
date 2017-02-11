package main

import(
	"os"
	"fmt"
)

func main() {
	
	kFlags.initializer(args, junkDesc)
	moveJunk()

}

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

//--------------------------------------------------- Dependencies: flagFilter, fmt
var argOne, argTwo, argThree, argFour string
// var fnOne, fnTwo, fnThree, fnFour func()

func Initializer(argsArray []string) {

	num := len(argsArray)

	if num > 5 {
                
		fmt.Println("Error: There should be no more than four arguments")
	} else if num == 5 {
                
		argOne, argTwo, argThree, argFour = os.Args[1], os.Args[2], os.Args[3], os.Args[4]
                
                flagFilter(m, argOne, argTwo, argThree, argFour)
	} else if num == 4 {
                
		argOne, argTwo, argThree, argFour = os.Args[1], os.Args[2], os.Args[3], ""
                
                flagFilter(m, argOne, argTwo, argThree)
	} else if num == 3 {
                
		argOne, argTwo, argThree, argFour = os.Args[1], os.Args[2], "", ""

                flagFilter(argOne, argTwo)                
	} else if num == 2 {
                
		argOne, argTwo, argThree, argFour = os.Args[1], "", "", ""

                flagFilter(m, argOne)
	} else {
                
		argOne, argTwo, argThree, argFour = "", "", "", ""

                fmt.Println(desc)
	}
}
//------------------------------------------------------------------------------------




//------------------------------------------------------------------------------------ Dependencies: blankFilter
var fn []func(){ fnOne, fnTwo, fnThree, fnFour}
fnOne, fnTwo, fnThree, fnFour := fn[0], fn[1], fn[2], fn[3]

func flagFilter(fMap map[string]func(), args ...string) ([]string) {
        
        blankFilter(args) 
        
        // iterates through and creates a byte slice for each argument 
        for x := 0; x <= len(args); x++ {
                var flag string
                argSlice := []byte(args[x])
                
                // iterates through the length of each arg slice in order to compare string to flags
                for i := 0; i <= len(argSlice); i++ {
                        var flagSlice []byte
                        
                        flagSlice[i] = argSlice[i] 
                        
                        // iterates through the flag/function map 
                        for _, k, v := range fMap {
                                if flagSlice == k {
                                        flag = k
                                        break
                                }
                        }
                        // checks whether flag has been set or not and then assigns the function based upon the flag
                        if flag != "" {
                               var param []byte
                               a := 0
                               
                               if flagSlice != "=" || flagSlice != " " {
                                       param[a] = flagSlice[i]
                                       a++
                               }
                               if i == argSlice {
                                        //would like to accept arguments here in version 2
                                        // fn[x] = ExecuteFlag(param)
                                        fn[x] = ExecuteFlag(flag)
                                        return fn[x]
                                       
                               }
                        }
                }
        }
}
//---------------------------------------------------------------------


// similar to BuildFlag, but it sets a function to be executed without a flag
//--------------------------------------------------------------------------- Dependencies: defaultFlagCheck
var iFnMap = make(map[string]func(int))
var sFnMap = make(map[string]func(string))

var intDefault bool
func DefaultFlagInt(i int, fn func(int)) {
        intDefault = true
        defaultFlagCheck()

}

var stringDefault bool
func DefaultFlagString(s string, fn func(string)) {
        stringDefault = true
        defaultFlagCheck()
        
}
//--------------------------------------------------------------------


// Throws an error if there is more than one default flag set
//---------------------------------------------------------------------------
func defaultFlagCheck(err error) error{
        
        if intDefault && stringDefault {
                return err
        } else {
                
        }
}
//-------------------------------------------------------------------------




//BuildFlag variables
// Builds flag for
//--------------------------------------------------------------------- Dependencies: N/A
var m = make(map[string]func())

func BuildFlag(flag string, fn func()) {
	
	m[flag] = fn
}
//---------------------------------------------------------------------




// Gets the function corresponding to the flag input
// not sure if I want this to be a public function
//-------------------------------------------------------------------- Dependencies: N/A
func ExecuteFlag(flag string) func(){

	fn := m[flag]
	return fn
}
//---------------------------------------------------------------------





// iterates through arguments to make sure none are empty
//-------------------------------------------------------------------- Dependencies: fmt
func blankFilter(s []string) {
        for _, x := range s {
                if x == "" {
                        fmt.Println(desc)
                }
        }
}
//-----------------------------------------------------------------------
