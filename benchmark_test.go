package goset

import "testing"

type Poem struct {
	Title string
	Body  []string
}

type Poet struct {
	Age   int
	Name  string
	Poems []*Poem
}

func IsSamePoet(a Poet, b Poet) bool {
	return a.Name == b.Name
}

func UniqPoets(poets []Poet) (r []Poet) {
	for _, p := range poets {
		found := false
		for _, added := range r {
			if IsSamePoet(p, added) {
				found = true
			}
		}
		if found {
			continue
		}
		r = append(r, p)
	}
	return
}

var LiBai = Poet{
	69,
	"李白",
	[]*Poem{
		&Poem{
			"望庐山瀑布",
			[]string{
				"日照香炉生紫烟",
				"遥看瀑布挂前川",
				"飞流直下三千尺",
				"疑是银河落九天",
			},
		},
		&Poem{
			"夜宿山寺",
			[]string{
				"危楼高百尺",
				"手可摘星辰",
				"不敢高声语",
				"恐惊天上人",
			},
		},
		&Poem{
			"关山月",
			[]string{
				"明月出天山",
				"苍茫云海间",
				"长风几万里",
				"吹度玉门关",
				"汉下白登道",
				"胡窥青海湾",
				"由来征战日",
				"不见有人还",
			},
		},
	},
}

var DuFu = Poet{
	57,
	"杜甫",
	[]*Poem{
		&Poem{
			"春望",
			[]string{
				"国破山河在",
				"城春草木深",
				"感时花溅泪",
				"恨别鸟惊心",
				"烽火连三月",
				"家书抵万金",
				"白头搔更短",
				"浑欲不胜簪",
			},
		},
		&Poem{
			"江南风李龟年",
			[]string{
				"岐王宅里寻常见",
				"崔九堂前几度闻",
				"正是江南好风景",
				"落花时节又逢君",
			},
		},
	},
}

var CuiHu = Poet{
	29,
	"崔护",
	[]*Poem{
		&Poem{
			"题南庄",
			[]string{
				"去年今日此门中",
				"人面桃花相映红",
				"人面不知何处去",
				"桃花依旧笑春风",
			},
		},
	},
}

var SuShi = Poet{
	40,
	"苏轼",
	[]*Poem{
		&Poem{
			"念奴娇",
			[]string{
				"大江东去浪",
				"淘尽千古风流",
				"人物故垒是三国",
				"周郎赤壁",
			},
		},
		&Poem{
			"东坡肉",
			[]string{
				"东坡肉真的好难吃",
				"东坡肉真的好难吃",
				"东坡肉真的好难吃",
				"东坡肉真的好难吃",
			},
		},
	},
}

var XinQiji = Poet{
	52,
	"辛弃疾",
	[]*Poem{
		&Poem{
			"郁孤台",
			[]string{
				"郁孤台下清江水",
				"中间多少行人泪",
				"西北望长安",
				"可怜无数山",
			},
		},
		&Poem{
			"破阵子",
			[]string{
				"鞋儿破",
				"帽儿破",
				"身上的袈裟破",
			},
		},
	},
}

func BenchmarkDiffernceBuiltinTypes(b *testing.B) {
	s1 := []string{"a", "b", "c", "d", "e", "f"}
	s2 := []string{"e", "f", "g", "h", "i", "j"}
	for i := 0; i < b.N; i++ {
		Difference(s1, s2)
	}
	return
}

func BenchmarkDiffernceCustomTypes(b *testing.B) {
	s1 := []Poet{LiBai, DuFu, CuiHu, SuShi}
	s2 := []Poet{LiBai, SuShi, XinQiji}

	for i := 0; i < b.N; i++ {
		Difference(s1, s2)
	}
	return
}

func BenchmarkUniqBuiltinTypes(b *testing.B) {
	s1 := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
	for i := 0; i < b.N; i++ {
		Uniq(s1)
	}
	return
}

func BenchmarkUniqCustomTypes(b *testing.B) {
	s1 := []Poet{LiBai, DuFu, CuiHu, SuShi, LiBai, SuShi, XinQiji}
	for i := 0; i < b.N; i++ {
		Uniq(s1)
	}
	return
}

func BenchmarkUniqPoets(b *testing.B) {
	s1 := []Poet{LiBai, DuFu, CuiHu, SuShi, LiBai, SuShi, XinQiji}
	for i := 0; i < b.N; i++ {
		UniqPoets(s1)
	}
	return
}
