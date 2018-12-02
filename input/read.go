package input

import (
	"io/ioutil"
)

func FromFile(path string) string {
	content, err :=	 ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(content[:])
}