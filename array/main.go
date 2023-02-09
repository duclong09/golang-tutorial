package main

import "fmt"

func main(){

	// points := [...]int{7,9,10,11,12}
	// fmt.Printf("%v, %T\n", points,points)

	
	numbers := [...]int{12, 78, 50}
	fmt.Println(numbers)


	listUser := [] int{2,4,6,8,10,12,14}
	//listProduct := listUser

	//ngan xep va hang doi:
	//Xoa cuoi: :6
	//listProduct := listUser[:5]
	listProduct := listUser[1:]
	//fmt.Println(listUser)
	fmt.Println(listProduct)
}