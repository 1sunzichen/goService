package main

import (
	"fmt"
)
//指针的零值
func nil1(){
		a:=25
		var b *int
		if b==nil{
			fmt.Println("b is",b)
			b=&a
			fmt.Println("b after initialization is",b)
		}
	
}
//使用新函数创建指针
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
//函数返回指针
func hello() *int{
	i:=5
	return &i
}
//不要将指向数组的指针作为参数传递给函数。改用切片。
func arrayFunc(arr *[3]int){
	(*arr)[0]=90
}
//尽管这种将指向数组的指针作为参数传递给函数并对其进行修改的方式有效，但这并不是 Go 中实现此目的的惯用方式。我们有切片。
func arrayFunc1(){
	a := [3]int{89, 90, 91}
    arrayFunc(&a)
    fmt.Println(a)
}
//基于同一个数组或切片创建的不同切片都共享同一个底层数组。如果一个切片修改了该底层数组的共享部分，其他切片和原始数组或切片都能感知到。其底层数据结构如下面两个图所示：
//共享同一底层数组
//切片 slices
func sliceFunc(sls []int) {  
    sls[0] = 90
}
//切片 slices2
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