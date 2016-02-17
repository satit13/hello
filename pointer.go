package main
import "fmt"

func main() {
	var p *int
	var i int
	var pn *int


	pn = new(int)

	*pn = 20  // assign value in pointer memory address "pn"

	fmt.Println(p)     // show memory address value  of "p"
	fmt.Println(i)     // show value of var "i"
	fmt.Println(&i)    // show memory address of var "i"
	fmt.Println(*pn)   // print value in memory address

}
