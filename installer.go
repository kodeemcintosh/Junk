package installer

import(
	"os"
	"fmt"
)

func main() {

}



func Init() {

	if runtime.GOOS == "windows" {
		fmt.Println("Hello from Windows")
	}	
}

