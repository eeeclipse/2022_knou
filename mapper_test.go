package main

import (
	"fmt"
)

// Creating Maps Using var and :=
// var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
// b := map[KeyType]ValueType{key1:value1, key2:value2,...}

// Creating Maps Using Using make()Function:
// var a = make(map[KeyType]ValueType)
// b := make(map[KeyType]ValueType)

// Creating an Empty Map
//var a map[KeyType]ValueType

func main() {
	var a = map[string]string{"brand": "Ford", "model": "mustang", "year": "1964"}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	var c = make(map[string]string) // The map is empty now
	c["brand"] = "Ford"
	c["model"] = "Mustang"
	c["year"] = "1964"
	// a is no longer empty
	d := make(map[string]int)
	d["Oslo"] = 1
	d["Bergen"] = 2
	d["Trondheim"] = 3
	d["Stavanger"] = 4

	var e = make(map[string]string)
	var f map[string]string

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)
	fmt.Printf("c\t%v\n", c)
	fmt.Printf("d\t%v\n", d)
	fmt.Printf("e\t%v\n", e)
	fmt.Printf("f\t%v\n", f)
}

