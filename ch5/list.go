package main

import "fmt"

func main() {
	var slice = []int{1,2,3,4,5}

	fmt.Printf("orig:   %v\n", slice)

	var addOne = func(x int) int {
		return x + 1
	}
	fmt.Printf("fmap:   %v\n", fmap(slice, addOne))

	var odd = func(x int) bool {
		return x % 2 == 1
	}
	fmt.Printf("filter: %v\n", filter(slice, odd))

	var even = func(x int) bool {
		return x % 2 == 0
	}
	fmt.Printf("filter: %v\n", filter(slice, even))

	fmt.Printf("bind:   %v\n", bind(slice, dup))
}

func dup(n int) []int {
	res := make([]int, n)
	for i := 0; i<n; i++ {
		res[i] = n
	}
	return res
}

func fmap(ls []int, f func(i int) int) []int {
	var res = make([]int, len(ls), cap(ls))
	for i, val := range ls {
		res[i] = f(val)
	}
	return res
}

func filter(ls []int, p func(int) bool) []int {
	res := make([]int, len(ls))
	pos := 0
	for _, val := range ls {
		if (p(val)) {
			res[pos] = val
			pos ++
		}
	}
	return res[:pos]
}

func bind(ls []int, f func(i int) []int) []int {
	res := make([]int, 0)
	for _, l := range ls {
		res = appendInt(res, f(l))
	}
	return res
}

func appendInt(x []int, y []int) []int {
	var z []int
    zlen := len(x) + len(y)
    if zlen <= cap(x) {
    	z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2 * len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z,x)
	}
	copy(z[len(x):], y)

	return z
}