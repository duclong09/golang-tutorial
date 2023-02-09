package main


import "fmt"


func main(){
	a := 10
	b := 0
	c := 5
	
	fmt.Println(finMax(a,b,c))
	fmt.Println(finSum(1,2,3))



}


func finMax(a, b , c int) ( int){
	max := a

	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	return max
	
}


func finSum(a ...int)(int){
	sum := 0
	for _, v := range a{
		sum += v
	}
	return sum
}


func calc(a,b, int) int{
	result := 0
	
}