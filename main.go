package main

import (
	"fmt"
	"go-practice/functions"
	"log"
)

func main() {
	b:= "shivank singh"
	c:=[]byte(b)
	log.Print(c)
sum:=functions.Add(2,3)

	const(
		width=5
		lenght=5
		area=width*lenght
		
	)
	fmt.Print(area)
	fmt.Print(sum)
}


