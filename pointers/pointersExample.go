package pointers

import "fmt"

func main() {
	x := 1
	y := &x //  assigns Y to the address of X

	fmt.Println("The value at x is: ", x)
	fmt.Println("The address of x is: ", &x)
	fmt.Println("The value of y is: ", y)
}
