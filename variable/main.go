package main

import "fmt"
import "strconv"
//Global Scope:
var(
	numberInt int  = 10	
	helloString string = "hello"
	numberFloat float64 = 11.1
)

var Number int = 12

func main() {
	//Block Scope
	numberOfDayPerYear := 365

	//%v : value , %T: type
	fmt.Printf("%v, %T", numberOfDayPerYear, numberOfDayPerYear)

	fmt.Println()
	helloStringName := "Dragon Viet Nam"
	fmt.Println(helloStringName)	

	//Add Global scope
	fmt.Println()
	fmt.Println(numberInt)
	fmt.Println(helloString)
	fmt.Println(numberFloat)
	//Convert type:

	fmt.Println()
	var conVertType float32 = float32(numberOfDayPerYear)
	fmt.Printf("%v, %T",conVertType,conVertType)
	fmt.Println()
	var strItoa string  = strconv.Itoa(numberOfDayPerYear)
	fmt.Printf("%v, %T",strItoa,strItoa)
}
