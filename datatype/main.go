package main

import "fmt"

func main(){
	//Kieu du lieu nguyen thuy


	//Boolean: true / false
	booleanTest := true
	fmt.Printf("%v, %T", booleanTest,booleanTest)


	//Number: int, float, complex:

	//int
	var numberOne int8 = 1
	var numberTwo int8 = 9
	//float:
	var numberFloatOne float32 = 3.14
	var numberFloatTwo float32 = 7.21

	fmt.Println()
	fmt.Printf("%v, %T\n", numberOne - numberTwo, numberOne - numberTwo)
	fmt.Printf("%v, %T\n", numberOne + numberTwo, numberOne + numberTwo)
	fmt.Printf("%v, %T\n", numberOne * numberTwo, numberOne * numberTwo)
	fmt.Printf("%v, %T\n", numberOne / numberTwo, numberOne / numberTwo)
	fmt.Printf("%v, %T\n", numberOne % numberTwo, numberOne % numberTwo)


	fmt.Println()
	fmt.Printf("%v, %T\n", numberFloatOne - numberFloatTwo, numberFloatOne - numberFloatTwo)
	fmt.Printf("%v, %T\n", numberFloatOne + numberFloatTwo, numberFloatOne + numberFloatTwo)
	fmt.Printf("%v, %T\n", numberFloatOne * numberFloatTwo, numberFloatOne * numberFloatTwo)
	fmt.Printf("%v, %T\n", numberFloatOne / numberFloatTwo, numberFloatOne / numberFloatTwo)



	//Text: string , byte, rune

	var stringTypeA = "Hello World"
	var stringTypeB = "Golang"

	fmt.Printf("%v, %T\n", stringTypeA + " " + stringTypeB, stringTypeA + stringTypeB)
}