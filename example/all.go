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
	// Uniq detection
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
	// Equal detection
	ok := goset.IsEqual(dupAvatars, filteredAvatars)
	fmt.Println(ok)

	// IsIncluded detection
	ok = goset.IsIncluded(dupAvatars, Avatar{112, "Angg"})
	fmt.Println("IsIncluded:", ok)

	// IsSubset detection
	ok = goset.IsSubset(
		[]Avatar{
			Avatar{112, "Angg"},
			Avatar{70, "Roku"},
		},
		dupAvatars,
	)
	fmt.Println("IsSubset:", ok)

	// IsSuperset detection
	ok = goset.IsSuperset(
		dupAvatars,
		[]Avatar{
			Avatar{112, "Angg"},
			Avatar{70, "Roku"},
		},
	)
	fmt.Println("IsSuperset:", ok)

	fmt.Println("== Operations ==")
	// Set operations
	aSet := []int{1, 2, 3}
	bSet := []int{3, 4, 5}
	//Difference
	uSet, iSet, aDiff, bDiff := goset.Difference(aSet, bSet)
	fmt.Println("Set Diffence:")
	fmt.Println("Intersetion:", iSet)
	fmt.Println("Union:", uSet)
	fmt.Println("A Differnce:", aDiff)
	fmt.Println("B Differnce:", bDiff)
	//Intersetion
	iSet = goset.Intersect(aSet, bSet).([]int)
	fmt.Println("Set Intersection:", iSet)
	//Union
	uSet = goset.Union(aSet, bSet).([]int)
	fmt.Println("Set Union:", uSet)
	//AddElement
	aSet = goset.AddElement(aSet, 4).([]int)
	fmt.Println("After AddElement", aSet)
	//AddElements
	aSet = goset.AddElements(aSet, []int{5, 6, 7}).([]int)
	fmt.Println("After AddElements", aSet)
	//RemoveElement
	aSet = goset.RemoveElement(aSet, 4).([]int)
	fmt.Println("After RemoveElement", aSet)
	//RemoveElements
	aSet = goset.RemoveElements(aSet, []int{5, 6, 7}).([]int)
	fmt.Println("After RemoveElements", aSet)
}
