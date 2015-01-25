package gomeh

import (
	"fmt"
	"io/ioutil"
	"log"
)

func readKey() string {
	buf, err := ioutil.ReadFile("./apikey")
	if err != nil {
		log.Fatal(err)
	}
	return string(buf)
}

func ExampleGetMeh_output() {
	apikey := readKey()
	m := GetMeh(apikey)
	fmt.Println(m)
}
