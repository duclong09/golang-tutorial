package main

import (
	"fmt"
)

const (
	//Enum:
	// red = 0
	// yellow = 1
	// blue = 2
	// black = 3


	//iota:
	// red = iota // 0 -> 1 -> 2 -> 3
	// yellow
	// blue
	// black

	//
	// errorDefault = iota // 1 -> 2 -> 3 -> 4
	// red 
	// yellow
	// blue
	// black

	//
	_ = iota // 1 -> 2 -> 3 -> 4
	red 
	yellow
	blue
	black



)

func main(){
	//hang so:
	
	fmt.Printf("%v, %T\n", red , red)
	fmt.Printf("%v, %T\n", yellow , yellow)
	fmt.Printf("%v, %T\n", blue , blue)
	fmt.Printf("%v, %T\n", black , black)
}