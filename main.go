//higher order functions are functions which are taking a functions an argument or they are retunrning a function

package main

import (
	"fmt"
	"time"
)

//  func operation(a int ,b int , op func(a int,b int)int )int{
// 	return op(a,b)
//  }

// func add (a int ,b  int ) int{
// 	return  a+b;
// }
func greet(ch chan string){
	ch <- "hello this is shivank singh"
}
func task() {
	fmt.Println("Task started")
	time.Sleep(2 * time.Second) // wait 2 seconds
	fmt.Println("Task done")
}


func main() {


	// mp:= map[string]int{
	// 	 "age":24,
	// }

	//  result,err:=functions.Checknumber(-1)
 	//  if err!=nil{

	// 	fmt.Errorf("this is the error we are getting %w",err)
	//  }
	//  fmt.Print(result)


// 	var a interface{}=4
//  value,ok:=a.(int)
//   if ok{
// 	fmt.Print("this is an integer value ")
//   }else{
// 	print(value)
//   }


// ch := make(chan string)  // make a channel

// 	go greet(ch)             // run function in background

// 	msg := <-ch              // receive message
// 	fmt.Println(msg)

go task()                   // runs in background
	fmt.Println("Main finished") // runs immediately

	time.Sleep(3 * time.Second) // so goroutine gets time to finish
	
}
