package main
import "fmt"
func main() {

	var n int
	var d float32
	var b bool
	var s string
	var ns [3]int
	var bs []byte
	var m map[string]string
	var pn *int


	fmt.Println("int:", n)
	fmt.Println("float32:", d)
	fmt.Println("bool:", b )
	fmt.Println("string:", s)
	fmt.Println("[3]int:", ns)
	fmt.Println("map[string]string:", m , m == nil)
	fmt.Println("*int:", pn)
	fmt.Println("Slice byte : ", bs, bs == nil)


	pn = new(int)
	*pn = 10
	fmt.Println("pointer", pn , pn ==nil)




}
