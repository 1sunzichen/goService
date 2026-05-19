package main
// When to use pointer receivers vs value receivers
// Generally, pointer receivers are used when changes made to the receiver inside the method
// should be visible to the caller. Pointer receivers can also be used where copying a data
// structure is expensive. Consider a struct with many fields. Using this struct as a value
// receiver in a method would require copying the entire struct, which would be expensive.
// In this case, if a pointer receiver is used, the struct is not copied and only a pointer
// to it is used in the method.
// In all other cases, value receivers can be used.



// Value receivers in methods vs value parameters in functions
// This topic is most suitable for beginners. I'll try to explain it clearly 😀.

// When a function has a value parameter, it will only accept a value argument.

// When a method has a value receiver, it will accept both pointer and value receivers.
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
//To define a method on a type, the receiver type definition and the method definition should exist in the same package.
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
