package main
import "fmt"

func main()  {
	fmt.Println(calculateBill(4,7))
}

func calculateBill(price int,no int)int{
	var totalPrice=price*no
	return totalPrice
}