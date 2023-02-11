package main
import (
	"fmt"
)
func map1(){
	employeeSalary := map[string]int{
        "steve": 12000,
        "jamie": 15000,
    }
    newEmp := "jamie"
    value, ok := employeeSalary[newEmp]
    if ok == true {
        fmt.Println("Salary of", newEmp, "is", value)
        return
    }
    fmt.Println(newEmp, "not found")
}
func map2(){
	employeeSalary2:=map[string]int{
		"steve":12000,
		"jamie":15000,
		"mike":9000,
	}
	fmt.Println("Content is map2")
	for key,value:=range employeeSalary2{
		fmt.Printf("employ[%s]=%d\n",key,value)
	}
}
func deletemap2(){
	employeeSalary3:=map[string]int{
		"steve":12000,
		"jamie":15000,
		"mike":9000,
	}
	fmt.Println("map before deletion",employeeSalary3)
	delete(employeeSalary3,"steve")
	fmt.Println("map after deletion",employeeSalary3)
}
func yinyong(){
	//map 引用类型
	employeeSalary := map[string]int{
        "steve": 12000,
        "jamie": 15000,     
        "mike": 9000,
    }
    fmt.Println("Original employee salary", employeeSalary)
    modified := employeeSalary
    modified["mike"] = 18000
    fmt.Println("Employee salary changed", employeeSalary)
}
func yinshepingdeng(){
	// map1 := map[string]int{
    //     "one": 1,
    //     "two": 2,
    // }

    // map2 := map1

    // if map1 === map2 {
	// 	fmt.Print("xiangdeng")
    // }else{
	// 	fmt.Print("不相等")
	// }
}
func main(){
	// employeeSalary:=make(map[string]int)
	// fmt.Println(employeeSalary)
	// employeeSalary["steve"]=12000;
	// employeeSalary["jamie"]=15000
	// employeeSalary["mike"]=9000
	// fmt.Println()
	// map1()
	// deletemap2()
	yinyong()
}