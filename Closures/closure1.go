package main

import "fmt"

func closures() func() int{
	count:=0

	return func() int{
		count++;
		return count
	}
}

func main2() {
	counter1:=closures()        //instance 1
	counter2:=closures()        //instance 2

	fmt.Println("this is first time we are calling closures with first instance ",counter1())
	fmt.Println("this is second time we are calling closures with first instance ",counter1())
    fmt.Println("this is third time we are calling closures with first instance ",counter1())

	fmt.Println("this is first time we are calling closures with second instance ",counter2())
	fmt.Println("this is second time we are calling closures with second instance ",counter2())
}