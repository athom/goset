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
	ok := goset.IsSubset(a1, a)
	fmt.Println(ok)
	a2 := []int{1, 2, 3, 5}
	ok = goset.IsSubset(a2, a)
	fmt.Println(ok)

	b := []string{"1", "2", "3", "4"}
	b1 := []string{"1", "2", "3", "4"}
	b2 := []string{"1", "2", "3", "5"}
	ok = goset.IsSubset(b1, b)
	fmt.Println(ok)
	ok = goset.IsSubset(b2, b)
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
	ok = goset.IsSubset(c1, c)
	fmt.Println(ok)
	ok = goset.IsSubset(c2, c)
	fmt.Println(ok)
}
