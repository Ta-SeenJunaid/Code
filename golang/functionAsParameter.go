package main

import "fmt"

type fn func(int)

func myfn1(i int) {
	fmt.Printf("\ni is %v\n", i)
}

func myfn2(i int) {
	fmt.Printf("\ni is %v\n", i)
}

func test(f fn, val int) {
	f(val)
}

func main() {
	test(myfn1, 1)
	test(myfn2, 2)
}
