package main

import (
	"fmt"
	"github.com/WildEgor/container-data-structures/pkg/set"
)

func main() {
	uniqueLetters := set.New[string]("A", "A", "B", "B")

	for _, s := range uniqueLetters.Values() {
		fmt.Println(s)
	}
}
