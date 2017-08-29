package main

import "encoding/json"
import "fmt"
import "os"

// structs for encoding / decoding custom types
type Response1 struct {
	Page   int
	Fruits []string
}

// tags on struct field declarations to customize encoded JSON key names
type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	// encoding basic data types to JSON strings
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// slices and maps encode to JSON arrays and objects
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// automatically encode  custom data types - default use field names as  JSON keys
	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &Response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// decoding JSON data into Go values
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// provide a variable to put the decoded data:  map[string]interface{} - arbitrary data types
	var dat map[string]interface{}

	// decoding - check for errors.
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// cast values in the decoded map to their appropriate type
	num := dat["num"].(float64)
	fmt.Println(num)

	// Accessing nested data requires a series of casts.
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// decode JSON into custom data types --> type-safety, eliminating the need for type assertions
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// can also stream JSON encodings directly to os.Writers like os.Stdout or to HTTP response bodies.
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
