// json-tags.go
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	type location struct {
		// tags for json
		// possible also for xml:"latitude"...
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	}

	curiosity := location{-4.5895, 137.4417}

	// basic usage:
	// bytes, err := json.Marshal(curiosity)

	// pretty printing:
	bytes, err := json.MarshalIndent(curiosity, "", "  ")
	exitOnError(err)

	fmt.Println(string(bytes))
}

// exitOnError prints any arrors and exists.
func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
