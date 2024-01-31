package main

import (
	"fmt"
	"strings"
)

func main() {
	// create a map using make
	//m := make(map[string]int)

	// map literal is also fine
	m := map[string]int{}

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	type LatLng struct {
		Lat, Long float64
	}

	// notice the VS Code note with yellow line!
	var locations = map[string]LatLng{
		"Bell Labs": LatLng{
			40.68433, -74.39967,
		},
		"Google": LatLng{
			37.42202, -122.08408,
		},
	}

	fmt.Println(locations)

	fmt.Println(WordCount("hello there"))
}

func WordCount(s string) map[string]int {
	m := map[string]int{}

	// what happens if i loop over the string without spliting it?
	// what is a rune!? alias for uint32 (UTF-32)
	// what is a string then? []byte slice

	for _, c := range strings.Split(s, " ") {
		m[c] = m[c] + 1
	}
	return m

}
