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
	ok := goset.IsIncluded(a, 1)
	fmt.Println(ok)
	ok = goset.IsIncluded(a, 5)
	fmt.Println(ok)

	b := []string{"1", "2", "3", "4"}
	ok = goset.IsIncluded(b, "1")
	fmt.Println(ok)
	ok = goset.IsIncluded(b, "5")
	fmt.Println(ok)

	c := []Avatar{
		Avatar{112, "Angg"},
		Avatar{70, "Roku"},
		Avatar{230, "Kyoshi"},
	}
	ok = goset.IsIncluded(c, Avatar{112, "Angg"})
	fmt.Println(ok)
	ok = goset.IsIncluded(c, Avatar{33, "Kuruk"})
	fmt.Println(ok)
}
