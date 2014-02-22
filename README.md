# Go Set

Go Set Package is a simple library for set operations with generic type supported.

[![Build Status](https://api.travis-ci.org/athom/goset.png?branch=master)](https://travis-ci.org/athom/goset)


## Installation

```bash
	go get "github.com/athom/goset"
```

## Features

- **Generic**

  All Go builtin types and custom defined types are supported.
  Even slice of pointers!

- **Handy**

  One Line Style calling design, aims to be **developer friendly**. 
  Give not enough shit on the performance and mathmatical rigour.


```go
	a := goset.Uniq([]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}).([]int)
	b := goset.Uniq([]string{"1", "2", "2", "3", "3", "3", "4"}).([]string)

        type Avatar struct {
                Age  int
                Name string
        }
	avatars := []Avatar{
		Avatar{112, "Angg"},
		Avatar{70, "Roku"},
		Avatar{230, "Kyoshi"},
		Avatar{230, "Kyoshi"},
		Avatar{33, "Kuruk"},
		Avatar{33, "Kuruk"},
		Avatar{33, "Kuruk"},
	}
	filteredAvatars := goset.Uniq(avatars).([]Avatar)
```

## Useage

#### Detections

###### 1. IsUniq

```go
	a := []int{1, 2, 3, 4, 4, 2, 3, 3, 4, 4}
	ua := goset.Uniq(a).([]int)
```

###### 2. IsEqual

```go
	a := []int{1, 2, 3}
	b := []int{2, 1, 3}
	ok := goset.IsEqual(a, b)
```

###### 3. IsIncluded

```go
	a := []int{1, 2, 3, 4}
	ok := goset.IsIncluded(a, 1)
```

###### 4. IsSubset

```go
	a := []int{1, 2, 3, 4}
	a1 := []int{1, 2, 3}
	ok := goset.IsSubset(a1, a)
```

###### 5. IsSuperset

```go
	a := []int{1, 2, 3, 4}
	a1 := []int{1, 2, 3}
	ok := goset.IsSuperset(a, a1)
```


#### Operations

###### 1. Uniq

```go
	a := []int{1, 2, 3, 4, 4, 2, 3, 3, 4, 4}
	ua := goset.Uniq(a).([]int)
```

###### 2. Intersect 

```go
	a1 := []int{1, 2, 3, 4}
	b1 := []int{3, 4, 5, 6}
	c1 := goset.Intersect(a1, b1).([]int)
```

###### 3. Union

```go
	a1 := []int{1, 2, 3, 4}
	b1 := []int{3, 4, 5, 6}
	c1 := goset.Union(a1, b1).([]int)
```

###### 4. Difference

```go
	a1 := []int{1, 2, 3, 4}
	b1 := []int{3, 4, 5, 6}
	_, _, c1, d1 := goset.Difference(a1, b1)
```

###### 5. AddElement

```go
	a := []int{1, 2, 3, 4}
	a = goset.AddElement(a, 5).([]int)
```

###### 6. AddElements

```go
	a := []int{1, 2, 3, 4}
	a = goset.AddElements(a, []int{5, 6}).([]int)
```

###### 7. RemoveElement

```go
	a := []int{1, 2, 3, 4}
	a = goset.RemoveElement(a, 4).([]int)
```

###### 8. RemoveElements

```go
	a := []int{1, 2, 3, 4}
	a = goset.RemoveElements(a, []int{3, 4}).([]int)
```



## License

Go Set is released under the [WTFPL License](http://www.wtfpl.net/txt/copying).


[![Bitdeli Badge](https://d2weczhvl823v0.cloudfront.net/athom/goset/trend.png)](https://bitdeli.com/free "Bitdeli Badge")

