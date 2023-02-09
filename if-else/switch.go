package main
import "fmt"


func main(){


	// number := 2
	// switch number{
	// 	case 1:
	// 		fmt.Println("one")
	// 	case 2:
	// 		fmt.Println("two")
	// 	case 3:
	// 		fmt.Println("three")

	// 	default:
	// 		fmt.Println("no order!")
	// }

	var i interface{} = "long"
	switch i.(type){
		case int:
			fmt.Println("type int")
		case string:	
			fmt.Println("type string")
		case float64:
			fmt.Println("type float")
		default : 
			fmt.Println("Undefine!")
	}
}