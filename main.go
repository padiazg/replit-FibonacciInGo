package main

import "fmt"

// 0 1 1 2 3 5 8 13 21, 34
/*  a0 a1 r0
0 :  0  0  0
1 :  1  0  1
2 :  0  1  1
3 :  1  1  2
4 :  1  2  3
5 :  2  3  5
6 :  3  5  8
7 :  5  8 13
8 :  8 13 21
*/

func fibonacci(a int) int {
	var a0, a1 int
	var r0 int

	for z := 0; z <= a; z++ {
		if z >= 2 {
			a0 = a1 // 
			a1 = r0
		} else {
			a0 = z
			a1 = 0
		}
		r0 = a0 + a1
		fmt.Printf("%v => (%v + %v) = %v\n", z, a0, a1, r0)
	}

	return r0
}

func main() {
	v0 := 50
	r0 := fibonacci(v0)
	fmt.Printf("fib(%v)=%v\n", v0, r0)
}
