package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func swap(a *int ,b *int){

	temp := *a
	*a=*b
	*b=temp
}
func main() {
	reader := bufio.NewReader(os.Stdin)
     
	fmt.Println("User please enter 1st number:")
	firstNumber, _ := reader.ReadString('\n')
	firstnumval, err := strconv.Atoi(strings.TrimSpace(firstNumber)) // Remove newline character
	if err != nil {
		fmt.Println("Invalid Integer:", err)
		return
	}

	 fmt.Println("User please enter 2nd number")
	 secondnumber,_:=reader.ReadString('\n')
	 secondnumval,err:=strconv.Atoi(strings.TrimSpace(secondnumber))
	 if err!=nil{
		fmt.Println("Invalid Integer",err)
		return
	 }
	 

	 swap(&firstnumval,&secondnumval)

	 fmt.Printf("Two numbers values after swap are :: First is %d Second is %d",firstnumval,secondnumval)

}