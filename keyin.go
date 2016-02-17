package main
import "fmt"
func main() {
	var i int
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Println(err)
	}else {
		switch {
		case i%5 ==0 && i%3 ==0 :
			fmt.Println("%3 , %5")
		case i%3 == 0 :
			fmt.Println("%3")
		case i%5 ==0 :
			fmt.Println("%5")
		default :
			fmt.Println("% Noting ")
		}
	}
}
