package main

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cast"
)

func main() {
	m1 := map[string]string{
		"name":   "kesa",
		"gender": "male",
	}

	m2 := map[string]interface{}{
		"name":   "kesa",
		"gender": "male",
	}

	m3 := map[interface{}]string{
		"name": "miao",
		"age":  "10",
	}

	m4 := map[interface{}]interface{}{
		"name": "miao",
		"age":  25,
	}

	jsonStr := `{"name":"pp","age": ""}`

	// ToStringMapString
	fmt.Println(cast.ToStringMapString(m1))      // map[gender:male name:kesa]
	fmt.Println(cast.ToStringMapString(m2))      // map[gender:male name:kesa]
	fmt.Println(cast.ToStringMapString(m3))      // map[age:10 name:miao]
	fmt.Println(cast.ToStringMapString(m4))      // map[age:25 name:miao]
	fmt.Println(cast.ToStringMapString(jsonStr)) // map[age: name:pp]

	m5 := make(map[string]interface{})
	json.Unmarshal([]byte(jsonStr), &m5)
	fmt.Println(m5) // map[age:222 name:pp]
}
