package main

import(
	"fmt"
)

func sendmessage(messagechannel chan string){
	messagechannel <- "Hello from Go routine"
}

func main(){
	message:=make(chan string)

	go sendmessage(message)

	fmt.Println("Hello from main and",<-message)
}