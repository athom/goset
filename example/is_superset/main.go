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
	a1 := []int{1, 2, 3}
	ok := goset.IsSuperset(a, a1)
	fmt.Println(ok)
	a2 := []int{1, 2, 3, 5}
	ok = goset.IsSuperset(a, a2)
	fmt.Println(ok)

	b := []string{"1", "2", "3", "4"}
	b1 := []string{"1", "2", "3", "4"}
	b2 := []string{"1", "2", "3", "5"}
	ok = goset.IsSuperset(b, b1)
	fmt.Println(ok)
	ok = goset.IsSuperset(b, b2)
	fmt.Println(ok)

	c := []Avatar{
		Avatar{112, "Angg"},
		Avatar{70, "Roku"},
		Avatar{230, "Kyoshi"},
	}
	c1 := []Avatar{
		Avatar{112, "Angg"},
		Avatar{70, "Roku"},
	}
	c2 := []Avatar{
		Avatar{112, "Angg"},
		Avatar{70, "Roku"},
		Avatar{33, "Kuruk"},
	}
	ok = goset.IsSuperset(c, c1)
	fmt.Println(ok)
	ok = goset.IsSuperset(c, c2)
	fmt.Println(ok)
}
