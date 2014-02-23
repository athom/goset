package goset

import "reflect"

// Operations
func Uniq(elements interface{}) interface{} {
	v := reflect.ValueOf(elements)
	if !isAvailableSlice(v) {
		return elements
	}
	if v.Len() <= 1 {
		return elements
	}

	slim := reflect.MakeSlice(reflect.TypeOf(elements), 0, v.Cap())
	for i := 0; i < v.Len(); i++ {
		found := false
		for j := 0; j < slim.Len(); j++ {
			if reflect.DeepEqual(
				v.Index(i).Interface(),
				slim.Index(j).Interface(),
			) {
				found = true
			}
		}
		if found {
			continue
		}
		slim = reflect.Append(slim, v.Index(i))
	}

	return slim.Interface()
}

func Intersect(aSet interface{}, bSet interface{}) interface{} {
	_, iSet, _, _ := Difference(aSet, bSet)
	return iSet
}

func Union(aSet interface{}, bSet interface{}) interface{} {
	uSet, _, _, _ := Difference(aSet, bSet)
	return uSet
}

func Difference(aSet interface{}, bSet interface{}) (iUnion, iIntersection, iADifferenceSet, iBDifferenceSet interface{}) {
	av := reflect.ValueOf(aSet)
	bv := reflect.ValueOf(bSet)
	if !areAvailableSlices(av, bv) {
		panic("A set and B set should be slices and have the same type of elements")
	}

	var union = reflect.MakeSlice(reflect.TypeOf(aSet), 0, av.Cap()+bv.Cap())
	var intersection = reflect.MakeSlice(reflect.TypeOf(aSet), 0, av.Cap()+bv.Cap())
	var aDifferenceSet = reflect.MakeSlice(reflect.TypeOf(aSet), 0, av.Cap())
	var bDifferenceSet = reflect.MakeSlice(reflect.TypeOf(aSet), 0, bv.Cap())

	for i := 0; i < av.Len(); i++ {
		skip := false
		for j := 0; j < bv.Len(); j++ {
			if reflect.DeepEqual(
				bv.Index(j).Interface(),
				av.Index(i).Interface(),
			) {
				intersection = reflect.Append(intersection, bv.Index(j))
				skip = true
			}
		}
		if !skip {
			aDifferenceSet = reflect.Append(aDifferenceSet, av.Index(i))
		}
	}

	for i := 0; i < bv.Len(); i++ {
		skip := false
		for j := 0; j < intersection.Len(); j++ {
			if reflect.DeepEqual(
				intersection.Index(j).Interface(),
				bv.Index(i).Interface(),
			) {
				skip = true
			}
		}
		if !skip {
			bDifferenceSet = reflect.Append(bDifferenceSet, bv.Index(i))
		}
	}

	union = reflect.AppendSlice(aDifferenceSet, intersection)
	union = reflect.AppendSlice(union, bDifferenceSet)

	iUnion = union.Interface()
	iIntersection = intersection.Interface()
	iADifferenceSet = aDifferenceSet.Interface()
	iBDifferenceSet = bDifferenceSet.Interface()

	return iUnion, iIntersection, iADifferenceSet, iBDifferenceSet
}

func AddElement(set interface{}, e interface{}) interface{} {
	v := reflect.ValueOf(set)
	if v.Type().Elem() != reflect.TypeOf(e) {
		panic("Set and element are not the same type")
	}

	if !isAvailableSlice(v) {
		panic("Invalid Slice")
	}

	if !IsUniq(set) {
		panic("Set should be uniq")
	}

	ev := reflect.ValueOf(e)

	for i := 0; i < v.Len(); i++ {
		if reflect.DeepEqual(
			e,
			v.Index(i).Interface(),
		) {
			return set
		}
	}

	v = reflect.Append(v, ev)
	return v.Interface()
}

func AddElements(aSet interface{}, bSet interface{}) interface{} {
	av := reflect.ValueOf(aSet)
	bv := reflect.ValueOf(bSet)
	if !areAvailableSlices(av, bv) {
		panic("Invalid Slices")
	}

	for i := 0; i < bv.Len(); i++ {
		aSet = AddElement(aSet, bv.Index(i).Interface())
	}
	return Uniq(aSet)
}

func RemoveElement(set interface{}, e interface{}) interface{} {
	v := reflect.ValueOf(set)
	if !isAvailableSlice(v) {
		panic("Invalid Slice")
	}

	if v.Len() == 0 {
		return set
	}

	ev := reflect.ValueOf(e)
	if !ev.IsValid() {
		return set
	}

	for i := 0; i < v.Len(); i++ {
		if reflect.DeepEqual(
			e,
			v.Index(i).Interface(),
		) {
			v = reflect.AppendSlice(
				v.Slice(0, i),
				v.Slice(i+1, v.Len()),
			)
			return v.Interface()
		}
	}

	return set
}

func RemoveElements(aSet interface{}, bSet interface{}) interface{} {
	av := reflect.ValueOf(aSet)
	bv := reflect.ValueOf(bSet)
	if !areAvailableSlices(av, bv) {
		panic("Invalid Slices")
	}

	for i := 0; i < bv.Len(); i++ {
		aSet = RemoveElement(aSet, bv.Index(i).Interface())
	}
	return Uniq(aSet)
}

// Detections
func IsUniq(aSet interface{}) bool {
	v := reflect.ValueOf(aSet)
	if !isAvailableSlice(v) {
		return false
	}
	if v.Len() <= 1 {
		return true
	}
	ele := v.Index(0).Interface()
	others := reflect.MakeSlice(reflect.TypeOf(aSet), v.Len()-1, v.Cap())
	for i := 1; i < v.Len(); i++ {
		if reflect.DeepEqual(
			ele,
			v.Index(i).Interface(),
		) {
			return false
		}
		others = v.Slice(1, v.Len())
	}
	return IsUniq(others.Interface())
}

func IsEqual(aSet interface{}, bSet interface{}) bool {
	av := reflect.ValueOf(aSet)
	bv := reflect.ValueOf(bSet)
	if av.Len() != bv.Len() {
		return false
	}
	if av.Len() == 0 && bv.Len() == 0 {
		return true
	}
	if !areAvailableSlices(av, bv) {
		return false
	}

	aMap := make(map[int]bool)
	bMap := make(map[int]bool)

	hits := 0
	for i := 0; i < av.Len(); i++ {
		if aMap[i] {
			continue
		}
		found := false
		for j := 0; j < bv.Len(); j++ {
			if bMap[j] {
				continue
			}
			if reflect.DeepEqual(
				av.Index(i).Interface(),
				bv.Index(j).Interface(),
			) {
				aMap[i] = true
				bMap[j] = true
				hits += 1
				found = true
			}
		}
		if !found {
			return false
		}
	}

	return hits == av.Len() && hits == bv.Len()
}

func IsIncluded(set interface{}, ele interface{}) bool {
	ev := reflect.ValueOf(ele)
	if !ev.IsValid() {
		return true
	}
	v := reflect.ValueOf(set)
	if !isAvailableSlice(v) {
		return false
	}
	if v.Len() == 0 {
		return false
	}
	if reflect.TypeOf(ev).String() != reflect.TypeOf(v.Index(0)).String() {
		return false
	}

	for i := 0; i < v.Len(); i++ {
		if reflect.DeepEqual(
			v.Index(i).Interface(),
			ev.Interface(),
		) {
			return true
		}
	}

	return false
}

func IsSubset(subSet interface{}, superSet interface{}) bool {
	_, _, aSubSet, _ := Difference(subSet, superSet)
	return reflect.ValueOf(aSubSet).Len() == 0
}

func IsSuperset(subSet interface{}, superSet interface{}) bool {
	_, _, _, bSubSet := Difference(subSet, superSet)
	return reflect.ValueOf(bSubSet).Len() == 0
}
