package main

import (
	"errors"
	"fmt"
	"strconv"
)

func numberInput() (int, error) {
	fmt.Println("Select Number (1 ~ ) : ")

	var nstr string
	fmt.Scanf("%s", &nstr)
	num, err := strconv.ParseUint(nstr, 10, 64)
	if err != nil {
		return -1, errors.New("Wrong Number")
	}

	return int(num), nil
}

func Menu(data []string, f func(string) error) error {
	for i, v := range data {
		fmt.Println(i+1, ":", v)
	}

	var num int
	var err error

	num, err = numberInput()
	if err != nil {
		return err
	} else if num == 0 {
		return nil
	} else if len(data) <= num-1 {
		return errors.New("Overwhelm Number")
	}

	err = f(data[num-1])
	if err != nil {
		return nil
	}
	return nil
}
