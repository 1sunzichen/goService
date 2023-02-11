package main
import ("fmt")
//可变参数
func hello(a int,b ...int){
	 fmt.Print(a,b)
}
func find(num int,nums ...int){
	fmt.Printf("type os nums is %T\n",nums)
	found := false
    for i, v := range nums {
        if v == num {
            fmt.Println(num, "found at index", i, "in", nums)
            found = true
        }
    }
    if !found {
        fmt.Println(num, "not found in ", nums)
    }
    fmt.Printf("\n")
}
func main(){
	// hello(5,1,2,3,4,5)
	find(89, 89, 90, 95)
    find(45, 56, 67, 45, 90, 109)
    find(78, 38, 56, 98)
    find(87)
}