package respool

import (
	"io/ioutil"
	"log"
)

var Res map[string][]byte

func Read(filename string) []byte {
	if val, ok := Res[filename]; ok {
		return val
	} else {
		text, err := ioutil.ReadFile(filename)
		check(err)
		Res[filename] = text
		return text
	}
	return []byte("Nothing")
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
