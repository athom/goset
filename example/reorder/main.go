package main

import (
	"fmt"

	"github.com/athom/goset"
)

type Avatar struct {
	Age  int
	Name string
}

func main() {
	avatars := []Avatar{
		Avatar{112, "Angg"},
		Avatar{70, "Roku"},
		Avatar{230, "Kyoshi"},
	}

	ageOrder := []int{70, 112, 230}
	fmt.Println("before reordered: ", avatars)
	avatars = goset.Reorder(ageOrder, avatars, "Age").([]Avatar)
	fmt.Println("after reordered: ", avatars)
	// output:
	// before reordered:  [{112 Angg} {70 Roku} {230 Kyoshi}]
	// after reordered:  [{70 Roku} {112 Angg} {230 Kyoshi}]
}
