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
	// Uniq operation
	a := goset.Uniq([]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}).([]int)
	fmt.Println(a)
	b := goset.Uniq([]string{"1", "2", "2", "3", "3", "3", "4"}).([]string)
	fmt.Println(b)

	avatars := []Avatar{
		Avatar{112, "Angg"},
		Avatar{70, "Roku"},
		Avatar{230, "Kyoshi"},
		Avatar{230, "Kyoshi"},
		Avatar{33, "Kuruk"},
		Avatar{33, "Kuruk"},
		Avatar{33, "Kuruk"},
	}
	// Uniq detector
	fmt.Println(goset.IsUniq(avatars))
	filteredAvatars := goset.Uniq(avatars).([]Avatar)
	fmt.Println(filteredAvatars)
	fmt.Println(goset.IsUniq(filteredAvatars))

	dupAvatars := []Avatar{
		Avatar{112, "Angg"},
		Avatar{70, "Roku"},
		Avatar{230, "Kyoshi"},
		Avatar{33, "Kuruk"},
	}
	// Equal detector
	equal := goset.Equal(dupAvatars, filteredAvatars)
	fmt.Println(equal)

	// Set operations
	aSet := []int{1, 2, 3}
	bSet := []int{3, 4, 5}
	//Difference
	uSet, iSet, aDiff, bDiff, _ := goset.Difference(aSet, bSet)
	fmt.Println("Set Diffence:")
	fmt.Println("Intersetion:", iSet)
	fmt.Println("Union:", uSet)
	fmt.Println("A Differnce:", aDiff)
	fmt.Println("B Differnce:", bDiff)
	//Intersetion
	iSet = goset.Intersection(aSet, bSet).([]int)
	fmt.Println("Set Intersection:")
	fmt.Println(iSet)
	//Union
	uSet = goset.Union(aSet, bSet).([]int)
	fmt.Println("Set Union:")
	fmt.Println(uSet)
}
