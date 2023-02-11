package main
import (
	"fmt"
	"unicode/utf8"
)
func printBytes(s string) {  
    fmt.Printf("Bytes: ")
    for i := 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
}

func printChars(s string) {  
    fmt.Printf("Characters: ")
    for i := 0; i < len(s); i++ {
        fmt.Printf("%c ", s[i])
    }
}
func printBytes2(s string) {  
    fmt.Printf("Bytes: ")
    for i := 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
}

func printChars2(s string) {  
    fmt.Printf("Characters: ")
    runes := []rune(s)
    for i := 0; i < len(runes); i++ {
        fmt.Printf("%c ", runes[i])
    }
}

func RuneCountInString1(){
	word1:="Senor"
	fmt.Printf("String: %s\n",word1)
	fmt.Printf("Length: %d\n",utf8.RuneCountInString(word1))
	fmt.Printf("Number of bytes: %d \n",len(word1))
	fmt.Printf("\n")
	word2:="pets"
	fmt.Printf("String: %s\n",word2)
	fmt.Printf("Length: %d\n",utf8.RuneCountInString(word2))
	fmt.Printf("Number of bytes: %d \n",len(word2))
	
}
func main() {  
    // name := "Hello World"
    // fmt.Printf("String: %s\n", name)
    // // printChars(name)
    // printChars2(name)
    // fmt.Printf("\n")
    // // printBytes(name)
	// printBytes2(name)
	RuneCountInString1()
}