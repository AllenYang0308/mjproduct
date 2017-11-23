package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	dat, err := ioutil.ReadFile("./pattern.txt")
	if err != nil {
		panic(err)
	}
	ss := strings.Split(strings.Replace(string(dat), "\\ ", "", -1), "\n")

	for _, v := range ss {
		if v != "" {
			fmt.Println(v)
		}
	}
}
