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
	a := []int{1, 2, 3, 4, 4, 2, 3, 3, 4, 4}
	ua := goset.Uniq(a).([]int)
	fmt.Println("before uniq: ", a)
	fmt.Println("after uniq: ", ua)

	b := []string{"1", "2", "2", "3", "3", "3", "4", "4", "4", "4"}
	ub := goset.Uniq(b).([]string)
	fmt.Println("before uniq: ", b)
	fmt.Println("after uniq: ", ub)

	c := []Avatar{
		Avatar{112, "Angg"},
		Avatar{70, "Roku"},
		Avatar{230, "Kyoshi"},
		Avatar{33, "Kuruk"},
		Avatar{112, "Angg"},
		Avatar{70, "Roku"},
	}
	uc := goset.Uniq(c).([]Avatar)
	fmt.Println("before uniq: ", c)
	fmt.Println("after uniq: ", uc)
}
