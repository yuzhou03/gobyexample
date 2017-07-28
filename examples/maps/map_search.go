package main

import "fmt"

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo)

	personDB["123"] = PersonInfo{"123", "Tom", "Beijing"}
	personDB["1"] = PersonInfo{"1", "Jack", "Shanghai"}

	id := "123"
	person, ok := personDB[id]

	if ok {
		fmt.Println("find person", person.Name, " with ID ", id)
		fmt.Printf("find person %s with ID %s", person.Name, id)
	} else {
		fmt.Println("not found")
	}
}
