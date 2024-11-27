package panic_recover

import "testing"
import "fmt"
//import "os"
import "errors"

func TestPanicVxExit(t *testing.T) {
	// defer func() {
	// 	fmt.Println("Finally!")
	// }()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from ", err)
		}
	}()
	fmt.Println("Start")
	//os.Exit(-1)
	panic(errors.New("Something wrong!"))
}