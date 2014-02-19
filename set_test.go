package goset

import (
	"labix.org/v2/mgo/bson"
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type TestSet struct{}

var _ = Suite(&TestSet{})

func (s *TestSet) TestIsEqual(c *C) {
	Id1 := bson.NewObjectId()
	Id2 := bson.NewObjectId()
	Id3 := bson.NewObjectId()
	Id4 := bson.NewObjectId()
	Id5 := bson.NewObjectId()

	var aSet []bson.ObjectId
	var bSet []bson.ObjectId

	//Happy path
	aSet = []bson.ObjectId{
		Id1,
		Id2,
		Id3,
	}
	bSet = []bson.ObjectId{
		Id1,
		Id2,
		Id3,
	}
	c.Check(IsEqual(aSet, bSet), Equals, true)

	//Both empty
	aSet = []bson.ObjectId{}
	bSet = []bson.ObjectId{}
	c.Check(IsEqual(aSet, bSet), Equals, true)

	//A is not empty B is empty
	aSet = []bson.ObjectId{
		Id1,
		Id2,
		Id3,
	}
	bSet = []bson.ObjectId{}
	c.Check(IsEqual(aSet, bSet), Equals, false)

	//A is empty B is not empty
	aSet = []bson.ObjectId{}
	bSet = []bson.ObjectId{
		Id1,
		Id2,
		Id3,
	}
	c.Check(IsEqual(aSet, bSet), Equals, false)

	//A contains B
	aSet = []bson.ObjectId{
		Id1,
		Id2,
		Id3,
	}
	bSet = []bson.ObjectId{
		Id1,
		Id2,
	}
	c.Check(IsEqual(aSet, bSet), Equals, false)
	return

	//A is subset of B
	aSet = []bson.ObjectId{
		Id1,
		Id2,
	}
	bSet = []bson.ObjectId{
		Id1,
		Id2,
		Id3,
	}
	c.Check(IsEqual(aSet, bSet), Equals, false)

	//A and B has intersection
	aSet = []bson.ObjectId{
		Id1,
		Id2,
		Id3,
	}
	bSet = []bson.ObjectId{
		Id2,
		Id3,
		Id4,
	}
	c.Check(IsEqual(aSet, bSet), Equals, false)

	//A and B has no intersection
	aSet = []bson.ObjectId{
		Id1,
		Id2,
		Id3,
	}
	bSet = []bson.ObjectId{
		Id4,
		Id5,
	}
	c.Check(IsEqual(aSet, bSet), Equals, false)
}

func (s *TestSet) TestIsUniq(c *C) {
	Id1 := bson.NewObjectId()
	Id2 := bson.NewObjectId()
	Id3 := bson.NewObjectId()
	Id4 := bson.NewObjectId()
	var set []bson.ObjectId

	//Happy path
	set = []bson.ObjectId{
		Id1,
		Id2,
		Id3,
		Id4,
	}
	c.Check(IsUniq(set), Equals, true)

	//Empty slice
	set = []bson.ObjectId{}
	c.Check(IsUniq(set), Equals, true)

	//One element
	set = []bson.ObjectId{
		Id1,
	}
	c.Check(IsUniq(set), Equals, true)

	//With repeat element
	set = []bson.ObjectId{
		Id1,
		Id1,
		Id2,
	}
	c.Check(IsUniq(set), Equals, false)
}

func (s *TestSet) TestIsIncluded(c *C) {
	Id1 := bson.NewObjectId()
	Id2 := bson.NewObjectId()
	Id3 := bson.NewObjectId()
	Id4 := bson.NewObjectId()
	var set []bson.ObjectId

	//Happy path
	set = []bson.ObjectId{
		Id1,
		Id2,
		Id3,
	}
	c.Check(IsIncluded(set, nil), Equals, true)
	c.Check(IsIncluded(set, Id1), Equals, true)
	c.Check(IsIncluded(set, Id2), Equals, true)
	c.Check(IsIncluded(set, Id3), Equals, true)
	c.Check(IsIncluded(set, Id4), Equals, false)

	//Empty slice
	set = []bson.ObjectId{}
	c.Check(IsIncluded(set, nil), Equals, true)
	c.Check(IsIncluded(set, Id1), Equals, false)
}

func (s *TestSet) TestIsSubset(c *C) {
	Id1 := bson.NewObjectId()
	Id2 := bson.NewObjectId()
	Id3 := bson.NewObjectId()
	Id4 := bson.NewObjectId()
	Id5 := bson.NewObjectId()
	var superSet1, superSet2, subSet1, subSet2, subSet3, subSet4, subSet5 []bson.ObjectId

	//Happy path
	superSet1 = []bson.ObjectId{
		Id1,
		Id2,
		Id3,
	}
	superSet2 = []bson.ObjectId{}
	subSet1 = []bson.ObjectId{
		Id1,
		Id2,
	}
	subSet2 = []bson.ObjectId{}
	subSet3 = []bson.ObjectId{
		Id1,
		Id2,
		Id3,
	}
	subSet4 = []bson.ObjectId{
		Id3,
		Id4,
	}
	subSet5 = []bson.ObjectId{
		Id1,
		Id2,
		Id3,
		Id4,
		Id5,
	}

	c.Check(IsSubset(subSet1, superSet1), Equals, true)
	c.Check(IsSubset(subSet2, superSet1), Equals, true)
	c.Check(IsSubset(subSet3, superSet1), Equals, true)
	c.Check(IsSubset(subSet4, superSet1), Equals, false)
	c.Check(IsSubset(subSet5, superSet1), Equals, false)
	c.Check(IsSubset(subSet2, superSet2), Equals, true)
	c.Check(IsSubset(subSet4, superSet2), Equals, false)
}

func (s *TestSet) TestIsSuperset(c *C) {
	Id1 := bson.NewObjectId()
	Id2 := bson.NewObjectId()
	Id3 := bson.NewObjectId()
	Id4 := bson.NewObjectId()
	Id5 := bson.NewObjectId()
	var superSet1, superSet2, subSet1, subSet2, subSet3, subSet4, subSet5 []bson.ObjectId

	//Happy path
	superSet1 = []bson.ObjectId{
		Id1,
		Id2,
		Id3,
	}
	superSet2 = []bson.ObjectId{}
	subSet1 = []bson.ObjectId{
		Id1,
		Id2,
	}
	subSet2 = []bson.ObjectId{}
	subSet3 = []bson.ObjectId{
		Id1,
		Id2,
		Id3,
	}
	subSet4 = []bson.ObjectId{
		Id3,
		Id4,
	}
	subSet5 = []bson.ObjectId{
		Id1,
		Id2,
		Id3,
		Id4,
		Id5,
	}

	c.Check(IsSuperset(superSet1, subSet1), Equals, true)
	c.Check(IsSuperset(superSet1, subSet2), Equals, true)
	c.Check(IsSuperset(superSet1, subSet3), Equals, true)
	c.Check(IsSuperset(superSet1, subSet4), Equals, false)
	c.Check(IsSuperset(superSet1, subSet5), Equals, false)
	c.Check(IsSuperset(superSet2, subSet2), Equals, true)
	c.Check(IsSuperset(superSet2, subSet4), Equals, false)
}

func (s *TestSet) TestUniq(c *C) {
	Id1 := bson.NewObjectId()
	Id2 := bson.NewObjectId()
	Id3 := bson.NewObjectId()
	Id4 := bson.NewObjectId()
	var set, eSet []bson.ObjectId

	//Happy path
	set = []bson.ObjectId{
		Id1,
		Id2,
		Id2,
		Id3,
		Id3,
		Id3,
		Id4,
		Id4,
		Id4,
		Id4,
	}
	eSet = []bson.ObjectId{
		Id1,
		Id2,
		Id3,
		Id4,
	}
	c.Check(IsEqual(eSet, Uniq(set).([]bson.ObjectId)), Equals, true)

	//Empty slice
	set = []bson.ObjectId{}
	eSet = []bson.ObjectId{}
	c.Check(IsEqual(eSet, Uniq(set).([]bson.ObjectId)), Equals, true)

	//One element
	set = []bson.ObjectId{
		Id1,
	}
	eSet = []bson.ObjectId{
		Id1,
	}
	c.Check(IsEqual(eSet, Uniq(set).([]bson.ObjectId)), Equals, true)
}

func (s *TestSet) TestDiffSet(c *C) {
	aId1 := bson.NewObjectId()
	aId2 := bson.NewObjectId()
	aId3 := bson.NewObjectId()
	aId4 := bson.NewObjectId()

	//A, B has intersection
	aSet := []bson.ObjectId{
		aId1,
		aId2,
		aId3,
		aId4,
	}

	bId1 := aId3
	bId2 := aId4
	bId3 := bson.NewObjectId()
	bId4 := bson.NewObjectId()
	bSet := []bson.ObjectId{
		bId1,
		bId2,
		bId3,
		bId4,
	}

	eUSet := []bson.ObjectId{
		aId1,
		aId2,
		aId3,
		aId4,
		bId3,
		bId4,
	}
	eISet := []bson.ObjectId{
		bId1,
		bId2,
	}

	eASet := []bson.ObjectId{
		aId1,
		aId2,
	}
	eBSet := []bson.ObjectId{
		bId3,
		bId4,
	}

	uSet, iSet, aSubSet, bSubSet := Difference(aSet, bSet)
	checkSet(eUSet, eISet, eASet, eBSet, uSet, iSet, aSubSet, bSubSet, c)

	//A contains B
	aSet = []bson.ObjectId{
		aId1,
		aId2,
		aId3,
		aId4,
	}

	bSet = []bson.ObjectId{
		aId3,
		aId4,
	}

	eUSet = []bson.ObjectId{
		aId1,
		aId2,
		aId3,
		aId4,
	}
	eISet = []bson.ObjectId{
		aId3,
		aId4,
	}
	eASet = []bson.ObjectId{
		aId1,
		aId2,
	}
	eBSet = []bson.ObjectId{}

	uSet, iSet, aSubSet, bSubSet = Difference(aSet, bSet)
	checkSet(eUSet, eISet, eASet, eBSet, uSet, iSet, aSubSet, bSubSet, c)

	//A inside B
	aSet = []bson.ObjectId{
		aId3,
		aId4,
	}

	bSet = []bson.ObjectId{
		aId1,
		aId2,
		aId3,
		aId4,
	}
	eUSet = []bson.ObjectId{
		aId1,
		aId2,
		aId3,
		aId4,
	}
	eISet = []bson.ObjectId{
		aId3,
		aId4,
	}
	eASet = []bson.ObjectId{}
	eBSet = []bson.ObjectId{
		aId1,
		aId2,
	}

	uSet, iSet, aSubSet, bSubSet = Difference(aSet, bSet)
	checkSet(eUSet, eISet, eASet, eBSet, uSet, iSet, aSubSet, bSubSet, c)

	//A and B has no commen
	aSet = []bson.ObjectId{
		aId1,
		aId2,
	}
	bSet = []bson.ObjectId{
		bId3,
		bId4,
	}

	eUSet = []bson.ObjectId{
		aId1,
		aId2,
		bId3,
		bId4,
	}
	eISet = []bson.ObjectId{}
	eASet = []bson.ObjectId{
		aId1,
		aId2,
	}
	eBSet = []bson.ObjectId{
		bId3,
		bId4,
	}

	uSet, iSet, aSubSet, bSubSet = Difference(aSet, bSet)
	checkSet(eUSet, eISet, eASet, eBSet, uSet, iSet, aSubSet, bSubSet, c)

	//String type related testing
	aStr1 := "a1"
	aStr2 := "a2"
	aStr3 := "a3"
	aStr4 := "a4"

	//A, B has intersection
	aStrSet := []string{
		aStr1,
		aStr2,
		aStr3,
		aStr4,
	}

	bStr1 := aStr3
	bStr2 := aStr4
	bStr3 := "b3"
	bStr4 := "b4"
	bStrSet := []string{
		bStr1,
		bStr2,
		bStr3,
		bStr4,
	}

	eStrUSet := []string{
		aStr1,
		aStr2,
		aStr3,
		aStr4,
		bStr3,
		bStr4,
	}
	eStrISet := []string{
		bStr1,
		bStr2,
	}

	eStrASet := []string{
		aStr1,
		aStr2,
	}
	eStrBSet := []string{
		bStr3,
		bStr4,
	}

	uStrSet, iStrSet, aStrSubSet, bStrSubSet := Difference(aStrSet, bStrSet)
	checkSetStr(eStrUSet, eStrISet, eStrASet, eStrBSet, uStrSet, iStrSet, aStrSubSet, bStrSubSet, c)

	//A contains B
	aStrSet = []string{
		aStr1,
		aStr2,
		aStr3,
		aStr4,
	}

	bStrSet = []string{
		aStr3,
		aStr4,
	}

	eStrUSet = []string{
		aStr1,
		aStr2,
		aStr3,
		aStr4,
	}
	eStrISet = []string{
		aStr3,
		aStr4,
	}
	eStrASet = []string{
		aStr1,
		aStr2,
	}
	eStrBSet = []string{}

	uStrSet, iStrSet, aStrSubSet, bStrSubSet = Difference(aStrSet, bStrSet)
	checkSetStr(eStrUSet, eStrISet, eStrASet, eStrBSet, uStrSet, iStrSet, aStrSubSet, bStrSubSet, c)

	//A inside B
	aStrSet = []string{
		aStr3,
		aStr4,
	}

	bStrSet = []string{
		aStr1,
		aStr2,
		aStr3,
		aStr4,
	}
	eStrUSet = []string{
		aStr1,
		aStr2,
		aStr3,
		aStr4,
	}
	eStrISet = []string{
		aStr3,
		aStr4,
	}
	eStrASet = []string{}
	eStrBSet = []string{
		aStr1,
		aStr2,
	}

	uStrSet, iStrSet, aStrSubSet, bStrSubSet = Difference(aStrSet, bStrSet)
	checkSetStr(eStrUSet, eStrISet, eStrASet, eStrBSet, uStrSet, iStrSet, aStrSubSet, bStrSubSet, c)

	//A and B has no commen
	aStrSet = []string{
		aStr1,
		aStr2,
	}
	bStrSet = []string{
		bStr3,
		bStr4,
	}

	eStrUSet = []string{
		aStr1,
		aStr2,
		bStr3,
		bStr4,
	}
	eStrISet = []string{}
	eStrASet = []string{
		aStr1,
		aStr2,
	}
	eStrBSet = []string{
		bStr3,
		bStr4,
	}

	uStrSet, iStrSet, aStrSubSet, bStrSubSet = Difference(aStrSet, bStrSet)
	checkSetStr(eStrUSet, eStrISet, eStrASet, eStrBSet, uStrSet, iStrSet, aStrSubSet, bStrSubSet, c)
}

func checkSet(eUSet, eISet, eASet, eBSet []bson.ObjectId, uSet, iSet, aSubSet, bSubSet interface{}, c *C) {
	assertSet(eUSet, uSet.([]bson.ObjectId), c)
	assertSet(eISet, iSet.([]bson.ObjectId), c)
	assertSet(eASet, aSubSet.([]bson.ObjectId), c)
	assertSet(eBSet, bSubSet.([]bson.ObjectId), c)
}

func checkSetStr(eStrUSet, eStrISet, eStrASet, eStrBSet []string, uStrSet, iStrSet, aStrSubSet, bStrSubSet interface{}, c *C) {
	assertSetStr(eStrUSet, uStrSet.([]string), c)
	assertSetStr(eStrISet, iStrSet.([]string), c)
	assertSetStr(eStrASet, aStrSubSet.([]string), c)
	assertSetStr(eStrBSet, bStrSubSet.([]string), c)
}

func assertSet(eS []bson.ObjectId, rS []bson.ObjectId, c *C) {
	c.Check(IsEqual(eS, rS), Equals, true)
}

func assertSetStr(eS []string, rS []string, c *C) {
	c.Check(IsEqual(eS, rS), Equals, true)
}
