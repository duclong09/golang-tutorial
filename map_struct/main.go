package main

import (
	"fmt"
	"reflect"
)



type User struct{
	id int
	name string
	isFemale bool
	age int
	subjects []string
}

type People struct{
	name string
	age int
}

type Product struct{
	People
	amount int `Maximum can't over 100`
	subjects []string
}

func main(){
	//map: key và type 
	studentNameAge := make(map[string]int)	
	studentNameAge = map[string]int{
		"Long" : 23,
		"Nhi" : 22,
	}
	copyMap := studentNameAge
	copyMap["Nhi"] = 23
	//studentNameAge["Yen"] = 2
	//delete(studentNameAge, "Yen")
	//_, isExits := studentNameAge["Long"]
	fmt.Println(studentNameAge)
	fmt.Println(copyMap)




	//struct:
	infoUser := User{
		id: 1,
		name: "Nhi",
		isFemale: true,
		age: 22,
		subjects: []string{"English", "Japan", "Computer"},
	}
	fmt.Println(infoUser.name)

	fmt.Println()
	//struct empty:
	sUser := User {}
	sUser.id = 1
	sUser.name = "Nhi"
	sUser.age = 22
	sUser.isFemale = true
	sUser.subjects = []string{"English", "Japan", "Computer"}

	fmt.Println(sUser)



	//Struct inline:

	student := struct {name string}{name: "Nhi"}
	//copy:
	copyStudent := &student
	copyStudent.name = "Long"
	fmt.Println(student)
	fmt.Println(copyStudent)




	pr := Product{}
	ref := reflect.TypeOf(pr)
	field, _ := ref.FieldByName("amount")
	fmt.Println(field)
	pr.name = "Nhi"
	pr.age = 22
	pr.amount = 2
	pr.subjects = []string{"Thuoc", "Son","Kem"}
	fmt.Println(pr)
}

