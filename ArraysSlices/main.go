package main

import (
	"fmt"
)
func subtactOne(numbers []int){
	for i:=range numbers{
		numbers[i]-=2
	}
}
func slice1(){
	a:=[5]int{76,43,56,79,90}
	var b []int=a[1:4]
	fmt.Println("b is",b)
}
func slice2(){
	c:=[]int{6,7,8}
	fmt.Println(c)
}
func slice4(){
	nos := []int{8, 7, 6}
    fmt.Println("slice before function call", nos)
    subtactOne(nos)                               //function modifies the slice
    fmt.Println("slice after function call", nos) //modifications are visible outside
}
func slice3(){
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
    fruitslice := fruitarray[1:3]
    fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice)) //length of is 2 and capacity is 6
    fruitslice = fruitslice[:4] //re-slicing furitslice till its capacity
    fmt.Println("After re-slicing length is",len(fruitslice), "and capacity is",cap(fruitslice))
}
// copy 切片之后 不再引用数组  可以对原始数组垃圾回收  内存优化
func countries()[]string{
	countries:=[]string{"USA","Singapore","Germany","India","Australia"}
	needCountries:=countries[:len(countries)-2]
	countrieCopy:=make([]string,len(needCountries))
	copy(countrieCopy,needCountries)
	return countrieCopy
}
func make1(){
	i:=make([]int,5,5)
	fmt.Println(i)
}
func slice5(){
	c:=countries()
	fmt.Println(c)
}
func arraygo(){
	a:=[...]string{"USA","China","India","Germany"}
	b:=a
	b[0]="China"
	fmt.Println("a is ",a)
	fmt.Println("b is",b)
}

func printarray(a [3][2]string) {  
    for _, v1 := range a {
        for _, v2 := range v1 {
            fmt.Printf("%s ", v2)
        }
        fmt.Printf("\n")
    }
}
func example(){
	a:=[3]int{12}
	fmt.Println(a)
}
func example1(){
	a:=[...]int{12,78,50}
	fmt.Println(a)
}
func main(){
	// a := [3][2]string{
    //     {"lion", "tiger"},
    //     {"cat", "dog"},
    //     {"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
    // }
    // printarray(a)
	slice5()
	// arraygo()
	// example1()
	// example()
	// var a [3]int
	// a[0]=12
	// a[1]=23
	// a[2]=122
	// fmt.Println(a)
}