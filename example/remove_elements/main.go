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
	fmt.Println("before remove_elements:(original set) ", a)
	q := goset.RemoveElements(a, []int{3, 4}).([]int)
	fmt.Println("after remove_elements: ", q)
	fmt.Println("after remove_elements(original set): ", a)

	b := []string{"1", "2", "3", "4"}
	fmt.Println("before remove_elements: ", b)
	b = goset.RemoveElements(b, []string{"3", "4"}).([]string)
	fmt.Println("after remove_elements: ", b)

	c := []Avatar{
		Avatar{112, "Angg"},
		Avatar{70, "Roku"},
		Avatar{230, "Kyoshi"},
		Avatar{33, "Kuruk"},
	}
	fmt.Println("before remove_elements: ", c)
	c = goset.RemoveElements(c, []Avatar{
		Avatar{230, "Kyoshi"},
		Avatar{33, "Kuruk"},
	}).([]Avatar)
	fmt.Println("after remove_elements: ", c)
}
