package main

import (
	"validator/rest"
)

func main() {
	// record, err := getCSVHeaders()
	// fmt.Println(record, err)

	// for idx, r := range record {
	// 	fmt.Println(idx, r)
	// }

	rest.Run()
}
