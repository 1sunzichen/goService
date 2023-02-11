package main

import "fmt"

type Vertex struct{
	x int
	y int
}
type Employee struct{
	firstname,lastname string
	age int
	salary int
}
//访问结构的各个字段
func structview(){
	em6:=Employee{
		firstname:"小",
		lastname:"明",
		age:55,
		salary:6000,
	}
	fmt.Println("FirstName",em6.firstname)
	fmt.Println("lastname",em6.lastname)
	fmt.Println("Salary",em6.salary)
	fmt.Printf("Salary: $%d\n", em6.salary)
	em6.salary = 6500
    fmt.Printf("New Salary: $%d", em6.salary)
}
type Employee0 struct {  
    firstName string
    lastName  string
    age       int
    salary    int
}
//结构体的零值
func struct0(){
	var emp4 Employee0 //zero valued struct
    fmt.Println("First Name:", emp4.firstName)
    fmt.Println("Last Name:", emp4.lastName)
    fmt.Println("Age:", emp4.age)
    fmt.Println("Salary:", emp4.salary)
}
func struct1(){
	emp5 := Employee0{
        firstName: "John",
        lastName:  "Paul",
    }
    fmt.Println("First Name:", emp5.firstName)
    fmt.Println("Last Name:", emp5.lastName)
    fmt.Println("Age:", emp5.age)
    fmt.Println("Salary:", emp5.salary)
}
//指向结构体的指针
func structPointer(){
	emp8:=&Employee0{
		firstName:"Sam",
		lastName:"Anderson",
		age:55,
		salary:6000,
	}
	fmt.Println("First Name",(*emp8).firstName)
	fmt.Println("Age",(*emp8).age)
}
type name struct {  
    firstName string
    lastName  string
}
//结构平等
func structEquality(){
	name1 := name{
        firstName: "Steve",
        lastName:  "Jobs",
    }
    name2 := name{
        firstName: "Steve",
        lastName:  "Jobs",
    }
    if name1 == name2 {
        fmt.Println("name1 and name2 are equal")
    } else {
        fmt.Println("name1 and name2 are not equal")
    }

    name3 := name{
        firstName: "Steve",
        lastName:  "Jobs",
    }
    name4 := name{
        firstName: "Steve",
    }

    if name3 == name4 {
        fmt.Println("name3 and name4 are equal")
    } else {
        fmt.Println("name3 and name4 are not equal")
    }
}
type image struct{
	data map[int]int

}
//如果结构变量包含不可比较的字段，则结构变量不可比较（感谢来自 reddit 的alasijia指出这一点）。
func structNotEqual(){
// image1 := image{
//         data: map[int]int{
//             0: 155,
//         }}
//     image2 := image{
//         data: map[int]int{
//             0: 155,
//         }}
//     if image1 == image2 {
//         fmt.Println("image1 and image2 are equal")
//     }
}
func main(){
	// v:=Vertex{1,2}
	// v.x=4
	// fmt.Println(v.x)
	// em1:=Employee{
	// 	firstname:"Sam",
	// 	age:25,
	// 	salary:500,
	// 	lastname:"Anderson",
	// }
	// em2:=Employee{"Thomas","Paul",29,800}
	// fmt.Println("Employee 1",em1)
	// fmt.Println("Employee 2",em2)
	// structview()
	// struct0()
	// struct1()
	// structPointer()
	// structEquality()
	structNotEqual()
}