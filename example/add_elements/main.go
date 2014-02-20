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
	a := []int{1, 2, 3, 4}
	fmt.Println("before add_elements: ", a)
	a = goset.AddElements(a, []int{5, 6}).([]int)
	fmt.Println("after add_elements: ", a)

	b := []string{"1", "2", "3", "4"}
	fmt.Println("before add_elements: ", b)
	b = goset.AddElements(b, []string{"5", "6"}).([]string)
	fmt.Println("after add_elements: ", b)

	c := []Avatar{
		Avatar{112, "Angg"},
		Avatar{70, "Roku"},
	}
	fmt.Println("before add_elements: ", c)
	c = goset.AddElements(c, []Avatar{
		Avatar{230, "Kyoshi"},
		Avatar{33, "Kuruk"},
	}).([]Avatar)
	fmt.Println("after add_elements: ", c)
}
