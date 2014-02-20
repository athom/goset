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
	fmt.Println("before differnce A:", a1)
	fmt.Println("before differnce B:", b1)
	_, _, c1, d1 := goset.Difference(a1, b1)
	fmt.Println("after differnce A': ", c1.([]int))
	fmt.Println("after differnce B': ", d1.([]int))

	a2 := []string{"1", "2", "3", "4"}
	b2 := []string{"3", "4", "5", "6"}
	fmt.Println("before differnce A:", a2)
	fmt.Println("before differnce B:", b2)
	_, _, c2, d2 := goset.Difference(a2, b2)
	fmt.Println("after differnce A': ", c2.([]string))
	fmt.Println("after differnce B': ", d2.([]string))

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
	fmt.Println("before differnce A:", a3)
	fmt.Println("before differnce B:", b3)
	_, _, c3, d3 := goset.Difference(a3, b3)
	fmt.Println("after differnce A': ", c3.([]Avatar))
	fmt.Println("after differnce B': ", d3.([]Avatar))
}
