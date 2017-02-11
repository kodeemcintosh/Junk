//package flags
package main

import "fmt"

func main() {

	// FUNCTION TO FUNCTION-OBJECT INITIALIZER
	//  |||
	//  vvv
	// var a func()
	// a = func(){
	// 	fmt.Print("a")
	// }
	// a()

	var testFunc = func() {
		fmt.Print("Hey, it worked")
		a[0] = 55
		b[2] = "is a number"
		fmt.Print(a[0])
		fmt.Print(b[2])
	}
	SetFlag("-Flag", testFunc)
	x = GetFunc("-Flag")
	x()
}


var m = make(map[string]func())

//BuildFlag variables
// Builds flag for
func SetFlag(flag string, fn func()) {

	m[flag] = fn
}

// Gets the function corresponding to the flag input
func GetFunc(flag string) func(){

	fn := m[flag]
	return fn
}

// var argOne, argTwo, argThree, argFour string
// func Initializer(argsArray []string, desc string) {
//
// 	num := len(argsArray)
//
// 	if num > 5 {
// 		fmt.Println("Error: There should be no more than four arguments")
// 	} else if num == 5 {
// 		argOne, argTwo, argThree, argFour = os.Args[1], os.Args[2], os.Args[3], os.Args[4]
// 		fmt.Println(argOne,argTwo,argThree,argFour)
// 	} else if num == 4 {
// 		argOne, argTwo, argThree, argFour = os.Args[1], os.Args[2], os.Args[3], ""
// 		fmt.Println(argThree,argTwo,argOne)
// 	} else if num == 3 {
// 		argOne, argTwo, argThree, argFour = os.Args[1], os.Args[2], "", ""
// 		fmt.Println(argTwo,argOne)
// 	} else if num == 2 {
// 		argOne, argTwo, argThree, argFour = os.Args[1], "", "", ""
// 		fmt.Println(argOne)
// 	} else {
// 		argOne, argTwo, argThree, argFour = "", "", "", ""
// 		fmt.Println(desc)
// 	}
// }

//func _flagExecute(arg string) (string){

//}


func filterArg(arg string) (bool) {

	for k,v := range m {
		if arg == k
		return true
	}
	

}
