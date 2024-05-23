package main

import "fmt"

// Maps are used to store data values in key:values pairs.
// Each element in map is key:value pair.
// A map is an unordered and changeable collection that does not allow duplicates.
// The length of map is the number of its elements. You can find it using the len() function.
// The default values of a map is nil
// Maps hold references to an underlying hash table.

// var a = map[keyType]valueType{key: value, key: value, ...}

var a = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
}

// using make keyword
var b = make(map[string]int)

func customMap() {
	// fmt.Println("Before delete the map is", a)
	// delete(a, "two")
	// fmt.Println("After delete the map is", a)

	// ok, _ := a["one"]
	// fmt.Println(ok)

	for key, val := range a {
		fmt.Println(key, val)
	}
}
