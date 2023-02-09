package main


import "fmt"

type myStruct struct{
	number int
}
func main(){
	// var a int = 12
	// var b *int = &a
	// fmt.Println(a, *b)
	// a = 24
	// fmt.Println(a, *b)



	// a := []int{1,2,3}
	// b := a
	// fmt.Println(a,b)
	// a[1] = 9
	// a[2] = 10
	// fmt.Println(&a[1],&b[2])


	var a *myStruct
	fmt.Println(a)
	a = new(myStruct)
	fmt.Println(a)
	(*a).number = 22
	fmt.Println((*a).number)
}


