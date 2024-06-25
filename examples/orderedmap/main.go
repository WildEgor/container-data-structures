package main

import (
	"fmt"
	"github.com/WildEgor/container-data-structures/pkg/orderedmap"
)

func main() {
	om := orderedmap.New[string, string]()

	om.Set("foo", "bar")
	om.Set("bar", "baz")
	om.Set("coucou", "toi")

	fmt.Println(om.Get("foo"))          // => "bar", true
	fmt.Println(om.Get("i dont exist")) // => "", false

	// iterating pairs from oldest to newest:
	for pair := om.Oldest(); pair != nil; pair = pair.Next() {
		fmt.Printf("%s => %s\n", pair.Key, pair.Value)
	}

	// iterating over the 2 newest pairs:
	i := 0
	for pair := om.Newest(); pair != nil; pair = pair.Prev() {
		fmt.Printf("%s => %s\n", pair.Key, pair.Value)
		i++
		if i >= 2 {
			break
		}
	}
}
