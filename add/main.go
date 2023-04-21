package main

import "fmt"

func add(a int,b int)int{
	return a+b
}

func arrays()[6]int{
	 primes:=[6]int{2,3,4,5,11,23}
	return primes
}

func main(){
	fmt.Println("sum is",add(1,3),arrays())
}