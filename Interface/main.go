package main

import (
	"fmt"

)

//Go 中，接口是一组方法签名。当一个类型为接口中的所有方法提供定义时，就说它实现了接口

type vower interface{
	vowerFind()[]rune
}

type MyString string

func(m MyString)vowerFind()[]rune{
	var vowels []rune
	for _, rune := range m {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

type SalaryCalculator interface {
	CalculateSalary() myInt
}

type Permanent struct {
	empId    myInt
	basicpay myInt
	pf       myInt
}

type Contract struct {
	empId    myInt
	basicpay myInt
}
type myInt int
//salary of permanent employee is the sum of basic pay and pf
func (i myInt) CalculateSalary() myInt {
	return i
}
func (p Permanent) CalculateSalary() myInt {
	return p.basicpay + p.pf
}

//salary of contract employee is the basic pay alone
func (c Contract) CalculateSalary() myInt {
	return c.basicpay
}
func totalExpense(s []SalaryCalculator) {
	var expense=myInt(0)
	for _, v := range s {
		expense =expense+ v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%v", expense)
}
func interface1(){
	//pemp1 := Permanent{
	//	empId:    1,
	//	basicpay: 5000,
	//	pf:       20,
	//}
	//pemp2 := Permanent{
	//	empId:    2,
	//	basicpay: 6000,
	//	pf:       30,
	//}
	//cemp1 := Contract{
	//	empId:    3,
	//	basicpay: 3000,
	//}
	employees := []SalaryCalculator{myInt(2)}
	totalExpense(employees)
}
type Worker interface {
	Work()
}

type Person struct {
	name string
	age int
}

func (p Person) Work() {
	fmt.Println(p.name, "is working")
}

func describe(w Worker) {
	fmt.Printf("Interface type %T value %v\n", w, w)
}
func WorkerFunc(){
	p := Person{
		name: "Naveen",
	}
	var w Worker = p
	describe(w)
	w.Work()
}
//空接口  由于空接口有零个方法，所有类型都实现了空接口
func describe1(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}
//类型断言
func assert(i interface{}) {
	s := i.(int) //get the underlying int value from i
	fmt.Println(s)
}
func assertFunc(){
	var s interface{} = 56
	assert(s)
}
func describeFunc(){
	s := "Hello World"
	describe1(s)
	i := 55
	describe1(i)
	strt := struct {
		name string
	}{
		name: "Naveen R",
	}
	describe1(strt)
}

type Describer interface{
	Describe()
}

func (p Person) Describe(){
	fmt.Printf("%s is %d years old\n",p.name,p.age)
}
type Address struct{
	state string
	country string
}
func (a *Address) Describe(){
	fmt.Printf("State %s Country %s",a.state,a.country)
}
func interface21(){
	var d1 Describer
	p1:=Person{"Sam",23}
	d1=p1
	d1.Describe()
	p2:=Person{"zd",18}
	d1=&p2
	d1.Describe()

	var d2 Describer
	a:=Address{"swim","jilin"}
	d2=&a
	d2.Describe()
}
func main(){
	//a:=MyString("HelloWorld")
	//result:=a.vowerFind()
	//fmt.Printf("%c",result)
	//interface1()

	//WorkerFunc()
	//describeFunc()//空接口
	assertFunc()//断言
}