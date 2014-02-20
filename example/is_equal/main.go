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
	a := []int{1, 2, 3}
	b := []int{2, 1, 3}
	ok := goset.IsEqual(a, b)
	fmt.Println(ok)

	a1 := []string{"1", "2", "3"}
	b1 := []string{"2", "1", "3"}
	ok = goset.IsEqual(a1, b1)
	fmt.Println(ok)

	a2 := []Avatar{
		Avatar{112, "Angg"},
		Avatar{70, "Roku"},
		Avatar{230, "Kyoshi"},
		Avatar{33, "Kuruk"},
	}
	b2 := []Avatar{
		Avatar{230, "Kyoshi"},
		Avatar{112, "Angg"},
		Avatar{70, "Roku"},
		Avatar{33, "Kuruk"},
	}
	ok = goset.IsEqual(a2, b2)
	fmt.Println(ok)
}
