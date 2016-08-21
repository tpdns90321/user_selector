package main

import "fmt"

func main() {
	var names []string
	var err error
	if names, err = GetUserNames(); err != nil {
		fmt.Println(err)
	} else {
		err = Menu(names, UserChange)
		if err != nil {
			fmt.Println(err)
		}
	}
}
