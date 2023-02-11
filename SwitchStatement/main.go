package main
import (
	"fmt"
	"math/rand"
)

func number() int {  
	num := 15 * 5 
	return num
}
func randloopFunc(){
	randloop:
	for{
		switch i:=rand.Intn(100); {
		case i%2==0:
			fmt.Printf("Generated even number %d",i)
			break randloop
		}
	}
}
func fallthroughFunc(){
	switch num:=number();{
	case num<50:
		fmt.Printf("%d is less than 50\n",num)
		fallthrough
	case num<100:

		fmt.Printf("%d is less than 100\n",num)
		fallthrough
	case num<150:
		fmt.Printf("%d is lesser then 200",num)
	}
}
func main()  {
	randloopFunc()
	// fallthroughFunc()
	// finger:=8
	// fmt.Printf("Finger %d is",finger)
	// switch finger {
	// case 1:
	// 	fmt.Println("Thumb")
	// case 2:
	// 	fmt.Println("Index")
	// case 3:
	// 	fmt.Println("Middle")
	// case 4:
	// 	fmt.Println("Ring")
	// case 5:
	// 	fmt.Println("Pinky")
	// default:
	// 	fmt.Println("incorrect finger number")
	// }
}