package main

import (
	"fmt"
)
//Zero value of a pointer
func nil1(){
		a:=25
		var b *int
		if b==nil{
			fmt.Println("b is",b)
			b=&a
			fmt.Println("b after initialization is",b)
		}
	
}
//Creating a pointer using the new function
func funcPointer(){
	size:=new(int)
	fmt.Printf("Size value is %d,type is %T,address is %v\n",*size,size,size)
	*size=85
	fmt.Println("New Size value is",*size)
}
func cancelPointer(){
	b:=255
	a:=&b
	fmt.Println("address of b is",a)
}
//Function returning a pointer
func hello() *int{
	i:=5
	return &i
}
//Do not pass a pointer to an array as a parameter to a function. Use slices instead.
func arrayFunc(arr *[3]int){
	(*arr)[0]=90
}
//Although passing a pointer to an array to a function and modifying it works, this is not the idiomatic way to achieve this in Go. We have slices.
func arrayFunc1(){
	a := [3]int{89, 90, 91}
    arrayFunc(&a)
    fmt.Println(a)
}
//Different slices created from the same array or slice all share the same underlying array.
//If one slice modifies the shared portion of the underlying array, other slices and the original
//array or slice will all perceive the change. The underlying data structure is shown in the two diagrams below:
//Share the same underlying array
//Slice slices
func sliceFunc(sls []int) {  
    sls[0] = 90
}
//Slice slices2
func sliceFunc2() {  
    a := [3]int{89, 90, 91}
    sliceFunc(a[:])
	fmt.Println(a)

}
func main(){
	// b:=255
	// var a *int=&b
	// fmt.Printf("type of a is %T\n",a)
	// fmt.Println("address of b is",a)
	// nil1()
	// funcPointer()
	// d:=hello()
	// fmt.Println("Value of d",*d)
	// arrayFunc1()
	sliceFunc2()
}