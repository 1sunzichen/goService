package main

import "fmt"

func main(){
	num:=10
	if num%2==0{
		fmt.Println("the number",num,"is even")
		return 
	}
	fmt.Println("The number",num,"is odd")
}