package main
// 何时使用指针接收器，何时使用值接收器
// 通常，当在方法内部对接收器所做的更改应该对调用者可见时，可以使用指针接收器。
// 指针接收器也可用于复制数据结构成本高昂的地方。考虑一个具有许多字段的结构。
// 将此结构用作方法中的值接收器将需要复制整个结构，这将是昂贵的。在这种情况下，
// 如果使用指针接收器，则不会复制该结构，而在方法中仅使用指向它的指针。
// 在所有其他情况下，可以使用值接收器。



// 方法中的值接收器与函数中的值参数
// 这个话题旅行最适合新手。我会尽量说清楚😀。

// 当一个函数有一个值参数时，它将只接受一个值参数。

// 当一个方法有一个值接收器时，它将同时接受指针和值接收器
import (
	"fmt"
	"math"
)

type Employee struct{
	name string
	salary int
	currency string
}
func (e Employee) displaySalary(){
	fmt.Printf("Salary of %s is %s%d",e.name,e.currency,e.salary)
}
type R struct{
	length int
	width int
}
type L struct{
	radius float64
}
func (r R) Area()int{
	return r.length * r.width
}
func (c L) Area()float64{
	return math.Pi*c.radius*c.radius
}
type Employee1 struct {  
    name string
    age  int
}

/*
Method with value receiver  
*/
func (e Employee1) changeName(newName string) {  
    e.name = newName
}
//要在类型上定义方法，接收者类型的定义和方法的定义应该存在于同一个包中。
//func (a int)add(b int){
//
//}
type myInt int
func (a myInt)add(b myInt)myInt{
	return a+b
}
/*
Method with pointer receiver  
*/
func (e *Employee1) changeAge(newAge int) {  
    e.age = newAge
}

func main(){
	a:=myInt(6)
	count:=a.add(10)
	fmt.Printf("值为%d",count)
	// emp1:=Employee{
	// 	name:"Sam Adolf",
	// 	salary:5000,
	// 	currency:"$",
	// }
	// emp1.displaySalary()
	// r:=R{
	// 	length:10,
	// 	width:5,
	// }
	// c:=L{
	// 	radius:12,
	// }
	// fmt.Printf("Area of rectangle %d\n",r.Area())
	// fmt.Printf("Area of circle %f\n",c.Area())
	//e := Employee1{
    //    name: "Mark Andrew",
    //    age:  50,
    //}
    //fmt.Printf("Employee name before change: %s", e.name)
    //e.changeName("Michael Andrew")
    //fmt.Printf("\nEmployee name after change: %s", e.name)
	//
    //fmt.Printf("\n\nEmployee age before change: %d", e.age)
    //e.changeAge(51)
    //fmt.Printf("\nEmployee age after change: %d", e.age)
}
