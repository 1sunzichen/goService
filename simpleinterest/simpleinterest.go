package Simpleinterestr
import "fmt"
//Calculate calculates and returns the simple interest 
//for a principal p, rate of interest r for time duration t years
func Calculate(p float64, r float64, t float64) float64 {  
    interest := p * (r / 100) * t
    return interest
}

func init(){
    fmt.Println("Simpleinterestr包初始化完成")
}