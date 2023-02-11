package main

import (
	"fmt"
	"time"
)

func worker(id int){
	for i:=0;i<3;i++{
		fmt.Printf("work %d\n",id)
	}
}
func main(){
	worker(1)
	go worker(2)

	go func(name string) {
		fmt.Println(name)
	}("goroutine2")

	time.Sleep(50*time.Microsecond)


}
