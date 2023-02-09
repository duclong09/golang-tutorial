
package main

import "fmt"

type class struct {
	className string
	students  []student
}

type student struct {
	name   string
	rollNo int
	city   string
}

func main() {
	george := student{"George", 35, "New York"}
	john := student{"George", 25, "London"}

	students := []student{george, john}
	class := class{"firstA", students}

	fmt.Printf("class is %v\n", class)
}