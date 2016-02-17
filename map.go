package main
import (
	"fmt"
	"encoding/json"
)

func main() {
	var m []map[string]string


	reqJson := `[
	{"firstname": "satit", "lastname":"Chomwattana"},
	{"firstname": "panic", "lastname":"Tripet"}
	]`


	json.Unmarshal([]byte(reqJson),&m)
	for _,d := range m{
		fmt.Println(d["firstname"])
	}
}
