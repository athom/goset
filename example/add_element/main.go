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
	fmt.Println("before add_element: ", a)
	a = goset.AddElement(a, 5).([]int)
	fmt.Println("after add_element: ", a)

	b := []string{"1", "2", "3", "4"}
	fmt.Println("before add_element: ", b)
	b = goset.AddElement(b, "5").([]string)
	fmt.Println("after add_element: ", b)

	c := []Avatar{
		Avatar{112, "Angg"},
		Avatar{70, "Roku"},
		Avatar{230, "Kyoshi"},
	}
	fmt.Println("before add_element: ", c)
	c = goset.AddElement(c, Avatar{33, "Kuruk"}).([]Avatar)
	fmt.Println("after add_element: ", c)
}
