// +build OMIT

package playground

import (
	"fmt"
	"time"
)

var a, b = 1, 2

func multiops(x int64, y int64) (int64, int64, int64, int64) {
	var a = 0
	fmt.Println(fmt.Sprintf("In function %d", a))
	return x + y, x * y, x, y
}

func main() {
	fmt.Println(fmt.Sprintf("In main %d", a))
	sum, product, _, _ := multiops(42, 13)
	fmt.Println(fmt.Sprintf("In main after function %d", a))
	baseString := "of %d and %d : %d\n"
	fmt.Printf("Sum (%T) "+baseString, 42, 13, sum, sum)
	fmt.Printf("Product "+baseString, 42, 13, product)

	var i = 10
	var f float64 = float64(i)
	f = f + 1

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	today = today + 2222
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

}
