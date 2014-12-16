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
	avatars := []Avatar{
		Avatar{112, "Angg"},
		Avatar{70, "Roku"},
		Avatar{230, "Kyoshi"},
		Avatar{33, "Kuruk"},
		Avatar{112, "Angg"},
		Avatar{70, "Roku"},
	}
	names := goset.Map(avatars, func(avatar Avatar) string {
		return avatar.Name
	}, []string{}).([]string)
	fmt.Println("before map: ", avatars)
	fmt.Println("after map: ", names)
}
