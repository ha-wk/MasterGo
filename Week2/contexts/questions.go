package main

import (
	"context"
	"fmt"
	"time"
)

//1️⃣ Implement a function that runs a goroutine for a maximum of 4 seconds using context.WithTimeout().


func runfunc(ctx context.Context){

	for {
		select {
		case <-ctx.Done(): // If context is canceled
			fmt.Println("Task stopped:", ctx.Err())
			return
		default:
			fmt.Println("Processing...")
			time.Sleep(1 * time.Second) // Simulating work
		}
	}

}
func main() {
	ctx,_ := context.WithTimeout(context.Background(),5*time.Second)
	
    go runfunc(ctx)
	time.Sleep(6 * time.Second)
	fmt.Println("Main function excited")
}


//2️⃣ Create a worker goroutine that stops manually using context.WithCancel().

// func runcontext(ctx context.Context){
// 	for{
//       select{
// 	  case <- ctx.Done() :
//         fmt.Println("Task stopped",ctx.Err())
// 		return
// 	  default:
// 		fmt.Println("Processing...")
// 		time.Sleep(1*time.Second)
// 	  }

// 	}
// }

// func main() {
// 	ctx,cancel:=context.WithCancel(context.Background())

// 	go runcontext(ctx)
// 	time.Sleep(5*time.Second)
// 	cancel()
// 	time.Sleep(2*time.Second)
// 	fmt.Println("exiting main function")
// }


//3️⃣ Write a function that sets a deadline of 5 seconds, but stops earlier if manually canceled.


// func runcontext(ctx context.Context){
// 	for{
// 		select{
// 		case <-ctx.Done():
// 			fmt.Println("Task Over",ctx.Err())
// 			return
// 		default:
// 			fmt.Println("Processing...")
// 			time.Sleep(1*time.Second)	
// 		}
// 	}
// }

// func main() {
// 	deadline :=time.Now().Add(5*time.Second)
// 	ctx,cancel := context.WithDeadline(context.Background(),deadline)

// 	go runcontext(ctx)
// 	time.Sleep(3*time.Second)
// 	cancel()
// 	fmt.Println("Finishing Main Go routine")
// }





//4️⃣ Modify a function to use context.WithValue() to pass user credentials and retrieve them inside a goroutine.

// func runcontextwithvalue(ctx context.Context,){

// 	userID, ok := ctx.Value("UserId").(int) // Type assertion to ensure correct type
// 	if !ok {
// 		fmt.Println("No valid user ID found in context")
// 		return
// 	}
// 	for{
// 		select{
// 		case <-ctx.Done():
// 			fmt.Println("This Context is finished now")
// 			return
// 		default :
// 		    fmt.Println("This worker is still going for user with user id :",userID)
// 			time.Sleep(1*time.Second)	
// 		}
// 	}
// }

// func main() {
// 	ctx:=context.WithValue(context.Background(),"UserId",2345)
// 	go runcontextwithvalue(ctx)
// 	time.Sleep(5*time.Second)
// 	fmt.Println("Main function excited")
// }