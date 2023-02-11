package main

import "fmt"

func main(){
	a:=true
	b:=true
	fmt.Println("a",a,"b",b)
	c:=a&&b
	fmt.Println("c:",c)
	d:=a||b
	fmt.Println("d",d)
	i:=56
	j:=67.8
	sum:=i+int(j)
	// sum:=i+int(67.8)
	fmt.Println(sum)
	
}