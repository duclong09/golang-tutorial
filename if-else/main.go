package main

import "fmt"

func main(){
	number := 90
	guess := 80

	if guess < 10 || guess > 100 {
		fmt.Println(" < 10 || > 100")
	}
	if guess >= 10 &&  guess <= 100{
		if guess < number{
			fmt.Println(" number > guess  ")
		}
		if guess > number{
			fmt.Println(" number < guess  ")
		}
		if guess == number{
			fmt.Println("Good Play")
		}
	}
}