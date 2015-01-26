package gomeh

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
)

func readKey() string {
	f := "./apikey"
	var key string
	// If API key file exists
	if _, err := os.Stat(f); err == nil {
		// Read API key from file
		fmt.Println("Reading from file.")
		buf, err := ioutil.ReadFile(f)
		if err != nil {
			log.Fatal(err)
		}
		key = string(buf)
	} else {
		fmt.Println("Reading from file.")
		// Read API key from env (for travis)
		key = os.Getenv("meh_apikey")
	}
	// Trim the string to remove any whitespace or linebreaks
	fmt.Println(key)
	return strings.Trim(key, " \n")
}

func ExampleGetMeh_output() {
	apikey := readKey()
	m, err := GetMeh(apikey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)
}

func TestGetMeh(t *testing.T) {
	apikey := readKey()
	m, err := GetMeh(apikey)
	if err != nil {
		t.Error("Failed to retreive data from API.")
	}
	if len(m.Deal.Title) == 0 {
		t.Error("Missing deal.")
	}
}
