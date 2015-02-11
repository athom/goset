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
	fmt.Println("before remove_element: ", a)
	a = goset.RemoveElement(a, 4).([]int)
	fmt.Println("after remove_element: ", a)

	b := []string{"1", "2", "3", "4"}
	fmt.Println("before remove_element: ", b)
	b = goset.RemoveElement(b, "4").([]string)
	fmt.Println("after remove_element: ", b)

	c := []Avatar{
		Avatar{112, "Angg"},
		Avatar{70, "Roku"},
		Avatar{230, "Kyoshi"},
		Avatar{33, "Kuruk"},
	}
	fmt.Println("before remove_element: ", c)
	c = goset.RemoveElement(c, Avatar{33, "Kuruk"}).([]Avatar)
	fmt.Println("after remove_element: ", c)
}
