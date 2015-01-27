package gomeh

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

func readKey() string {
	f := "./apikey"
	var key string
	// If API key env exists use that
	if env := os.Getenv("meh_apikey"); env != "" {
		// Read API key from env (for travis)
		key = env
	} else {
		// Read API key from file, if exists
		if _, err := os.Stat(f); err == nil {
			buf, err := ioutil.ReadFile(f)
			if err != nil {
				log.Fatal(err)
			}
			key = string(buf)
		}
	}
	// Trim the string to remove any whitespace or linebreaks
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

func TestSoldOut(t *testing.T) {
	m := Meh{
		Deal: Deal{
			SoldOutAt: time.Now(),
		},
	}
	if m.SoldOut() != true {
		t.Error("Failed to correctly recognize as Sold Out.")
	}
}

func TestString(t *testing.T) {
	// Available product
	m := Meh{
		Deal: Deal{
			Title: "Product",
			Items: []Item{
				Item{Price: 5},
			},
		},
	}
	if fmt.Sprint(m) != fmt.Sprint("Product - $5") {
		t.Error("Failed to return available product string.")
	}
	// Sold out produc
	// Available product
	m = Meh{
		Deal: Deal{
			Title:     "Product2",
			SoldOutAt: time.Now(),
			Items: []Item{
				Item{Price: 5},
			},
		},
	}
	if fmt.Sprint(m) != fmt.Sprint("[Sold Out] Product2 - $5") {
		t.Error("Failed to return available product string.")
	}
}
