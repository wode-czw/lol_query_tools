package tools

import "fmt"

func Test_map() {

	var wode map[string]string
	wode = make(map[string]string, 20)

	wode["fizz"] = "like water"
	wode["akl"] = "like use fb"

	value, ok := wode["fzz"]
	if ok == true {
		fmt.Println("exist")
	}

	fmt.Println(value)
}
