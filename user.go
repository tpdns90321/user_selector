package main

import (
	"errors"
	"io/ioutil"
	"regexp"
)

func GetUserNames() ([]string, error) {
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

	return user, nil
}
