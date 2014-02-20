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
	ok := goset.IsUniq(a)
	fmt.Println(ok)
	a = []int{1, 2, 3, 4, 2}
	ok = goset.IsUniq(a)
	fmt.Println(ok)

	b := []string{"1", "2", "3", "4"}
	ok = goset.IsUniq(b)
	fmt.Println(ok)
	b = []string{"1", "2", "3", "4", "2"}
	ok = goset.IsUniq(b)
	fmt.Println(ok)

	c := []Avatar{
		Avatar{112, "Angg"},
		Avatar{70, "Roku"},
		Avatar{230, "Kyoshi"},
		Avatar{33, "Kuruk"},
	}
	ok = goset.IsUniq(c)
	fmt.Println(ok)

	c = []Avatar{
		Avatar{230, "Kyoshi"},
		Avatar{112, "Angg"},
		Avatar{70, "Roku"},
		Avatar{33, "Kuruk"},
		Avatar{230, "Kyoshi"},
	}
	ok = goset.IsUniq(c)
	fmt.Println(ok)
}
