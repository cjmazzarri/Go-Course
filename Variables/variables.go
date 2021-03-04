package main

import "fmt"

func main()  {
	var x, y int //Initial values 0, 0
	/*var word string 		Initial value ""
	var flag bool			Initial value false
	var str_arr []string 	Declaration for array of strings
	Variables cannot left unused*/
	x = 23
	y = 2
	z:= 15 //Dynamic typing, no need to specify type in declaration, detected by compiler
	//z:=16 error; can't use := again on a variable of the same name
	fmt.Println(x, y, z)
	fmt.Println(x + y)
	fmt.Println("x: ", x)
}