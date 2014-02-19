package goset

import (
	"errors"
	"reflect"
)

func isAvailableSlice(v reflect.Value) bool {
	if v.Kind() != reflect.Slice {
		return false
	}

	var kind string
	for i := 0; i < v.Len(); i++ {
		eleKind := reflect.TypeOf(v.Index(i)).Kind().String()
		if i == 0 {
			kind = eleKind
		} else {
			if kind != eleKind {
				return false
			}
		}
	}

	return true
}

func areAvailableSlices(v1, v2 reflect.Value) bool {
	if !isAvailableSlice(v1) {
		return false
	}
	if !isAvailableSlice(v2) {
		return false
	}
	if v1.Len() == 0 && v2.Len() != 0 {
		return false
	}
	if v1.Len() != 0 && v2.Len() == 0 {
		return false
	}
	if v1.Len() == 0 && v2.Len() == 0 {
		return true
	}
	return reflect.TypeOf(v1).Kind().String() == reflect.TypeOf(v2).Kind().String()
}

func Intersection(aSet interface{}, bSet interface{}) (iSet interface{}) {
	_, iSet, _, _, _ = Difference(aSet, bSet)
	return
}

func Union(aSet interface{}, bSet interface{}) (uSet interface{}) {
	uSet, _, _, _, _ = Difference(aSet, bSet)
	return
}

func Difference(aSet interface{}, bSet interface{}) (iUnion, iIntersection, iADifferenceSet, iBDifferenceSet interface{}, err error) {
	av := reflect.ValueOf(aSet)
	bv := reflect.ValueOf(bSet)
	if !areAvailableSlices(av, bv) {
		err = errors.New("A set and B set should be slices and have the same type of elements")
		return
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

	return
}

func IsUniq(aSet interface{}) (r bool) {
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

func Uniq(elements interface{}) (r interface{}) {
	v := reflect.ValueOf(elements)
	if !isAvailableSlice(v) {
		r = elements
		return
	}
	if v.Len() <= 1 {
		r = elements
		return
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

	r = slim.Interface()
	return

}

func Equal(aSet interface{}, bSet interface{}) (r bool) {
	av := reflect.ValueOf(aSet)
	bv := reflect.ValueOf(bSet)
	if av.Len() != bv.Len() {
		r = false
		return
	}
	if av.Len() == 0 && bv.Len() == 0 {
		r = true
		return
	}
	if !areAvailableSlices(av, bv) {
		r = false
		return
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
