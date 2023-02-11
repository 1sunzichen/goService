package main
import "fmt"
func Loops1(){
	for i := 0; i < 10; i++ {
		if i>5{
			break
		}
		fmt.Printf("%d ",i)
	}
	fmt.Printf("\nline after for loop")

}
func Loops2(){
	for i := 0; i < 10; i++ {
		if i%2==0 {
			continue
		}
		fmt.Printf("%d ",i)
	}
	fmt.Printf("\nline after for loop")

}
func Loops3(){
   
	for i := 1; i <= 3; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
	

}
func Loops4(){
	ppp:
	for i := 0; i <= 3; i++ {
		for j := 1; j <= 4; j++ {
			fmt.Printf("i = %d, j=%d \n",i,j)
			if(j==i){
				break ppp;
			}
		}
	}
}
func main()  {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n",i)
	}
	
	Loops4()
}