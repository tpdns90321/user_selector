package main

import "fmt"

func Example_changeuser() {
	UserChange("programming")
	//Output:
	//
}

func Example_getusername() {
	data, err := GetUserNames()
	fmt.Println(data, err)

	//Output:
	//programming <nil>
}
