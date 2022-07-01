package main

import "fmt"
 
/*
f(0)=0; f(1)=1; fn=f(n-1)-f(n-2)
0 1 1 2 3 5 8 13 21 34

    f0 f1 r0
0 :  0  0  0
1 :  0  1  1
2 :  0  1  1
3 :  1  1  2
4 :  1  2  3
5 :  2  3  5
6 :  3  5  8
7 :  5  8 13
8 :  8 13 21
*/

func fibonacci(a int) int {
	var (
    f0, f1, fn int
  )

  if a < 2 {
      return a
  } else {
    f1=0
    fn=1
    for z := 2; z <= a; z++ {
      f0 = f1
      f1 = fn
  		fn = f0 + f1
  		fmt.Printf("%v => (%v + %v) = %v\n", z, f0, f1, fn)
  	}  
  }

	return fn
}

func main() {
	v0 := 50
	r0 := fibonacci(v0)
	fmt.Printf("fib(%v)=%v\n", v0, r0)
}
