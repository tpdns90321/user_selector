package main

import (
	"errors"
	"io/ioutil"
	"regexp"
)

/*#include <stdio.h>
void command(const char *string){
	system(string);
}*/
import "C"

var usernames []string

func GetUserNames() ([]string, error) {
	if usernames == nil {
		f, err := ioutil.ReadFile("/etc/passwd")
		if err != nil {
			return nil, err
		}
		file := string(f)

		mat := regexp.MustCompile("/home/([\\d\\w/]+):/bin/[\\w]+")
		data := mat.FindAllStringSubmatch(file, -1)

		if data == nil {
			return nil, errors.New("No Available Users")
		}

		var user []string

		for _, v := range data {
			user = append(user, v[1])
		}

		usernames = user

		return user, nil
	}
	return usernames, nil
}

func UserChange(user string) error {
	ok := false
	users, _ := GetUserNames()
	for _, v := range users {
		if v == user {
			ok = true
		}
	}

	if !ok {
		return errors.New("No Available Users")
	}

	C.command(C.CString("su - " + user))
	return nil
}
