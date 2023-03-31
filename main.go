package main

import (
	"encoding/json"
	"fmt"
	"log"

	depsresolver "github.com/magdyamr542/dips/deps_resolver"
)

func main() {
	var resolver depsresolver.Resolver = depsresolver.NewResolver()

	entities := []string{"1", "2", "3", "4", "5", "6", "7"}
	depsStr := []byte(`{
		"1"  : ["2"],
		"2"  : ["3"],
		"3"  : ["4" , "5"],
		"5"  : ["4"],
		"6"  : ["7"],
		"7"  : ["5" , "3"]
	}`)
	deps := make(map[string][]string)
	err := json.Unmarshal(depsStr, &deps)
	if err != nil {
		log.Fatalf("cannot convert deps string to a map: %v\n", err)
	}

	order, err := resolver.Resolve(entities, deps)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Order is: %v\n", order)
}
