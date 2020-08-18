package main

import (
	"fmt"
	"sort"
)

func main() {

	// Make a map of index:string value:string
	states := make(map[string]string)
	fmt.Println(states)

	// set vlues in the map
	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CA"] = "California"
	states["AZ"] = "Arazona"
	states["NM"] = "New Mexico"
	states["NV"] = "Nevada"
	fmt.Println(states)

	california := states["CA"]
	fmt.Println(california)

	delete(states, "OR")
	fmt.Println(states)

	states["NY"] = "New York"

	for k, v := range states {
		fmt.Printf("%v : %v\n", k, v)
	}

	// make a slice of keys
	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println("\n\nSorted:")
	for i := range keys {
		println(states[keys[i]])
	}

}
