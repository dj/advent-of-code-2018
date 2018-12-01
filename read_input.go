package main

import (
	"io/ioutil"
)

func readInput(path string) string {
	content, err :=	 ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(content[:])
}