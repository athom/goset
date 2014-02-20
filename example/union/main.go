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
	a1 := []int{1, 2, 3, 4}
	b1 := []int{3, 4, 5, 6}
	fmt.Println("before union A:", a1)
	fmt.Println("before union B:", b1)
	c1 := goset.Union(a1, b1).([]int)
	fmt.Println("after union: ", c1)

	a2 := []string{"1", "2", "3", "4"}
	b2 := []string{"3", "4", "5", "6"}
	fmt.Println("before union A:", a2)
	fmt.Println("before union B:", b2)
	c2 := goset.Union(a2, b2).([]string)
	fmt.Println("after union: ", c2)

	a3 := []Avatar{
		Avatar{112, "Angg"},
		Avatar{70, "Roku"},
		Avatar{230, "Kyoshi"},
	}
	b3 := []Avatar{
		Avatar{70, "Roku"},
		Avatar{230, "Kyoshi"},
		Avatar{33, "Kuruk"},
	}
	fmt.Println("before union A:", a3)
	fmt.Println("before union B:", b3)
	c3 := goset.Union(a3, b3).([]Avatar)
	fmt.Println("after union: ", c3)
}
