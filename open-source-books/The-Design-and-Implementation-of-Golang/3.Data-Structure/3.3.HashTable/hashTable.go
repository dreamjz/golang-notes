package hashtable

import "fmt"

func mapInit() {
	// literal
	hashMap := map[string]string{
		"A": "a",
		"B": "b",
		"C": "c",
	}
	fmt.Println(hashMap)
	// runtime
	hm := make(map[string]int, 3)
	hm["1"] = 1
	hm["2"] = 2
	hm["3"] = 3
	fmt.Println(hm)
}

func accessMap() {
	hm := make(map[string]int)
	// insert
	hm["A"] = 1
	hm["B"] = 2
	hm["C"] = 3
	// read
	fmt.Println("Hm[\"A\"]: ", hm["A"])
	// update
	hm["A"] = 0
	fmt.Printf("Updated hm[\"A\"]: %v\n", hm["A"])
	// delete
	delete(hm, "A")
	fmt.Println("Deleted Key A: ", hm)
}
