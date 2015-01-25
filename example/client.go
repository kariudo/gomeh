// Simple Meh client example for gomeh
package main

import (
	"fmt"

	"github.com/kariudo/gomeh"
)

const apikey = "n8Gjouad3YHVM554OgotOEW6Z6arjei5"

func main() {
	m := gomeh.GetMeh(apikey)
	fmt.Println(m)
}
