package main

import (
	"fmt"
	"strconv"
)
//Reminder: Imported packages must be used 
func main()  {
	age := 23
	//fmt.Println("My age is"+age) would result in an error due to types
	age_str := strconv.Itoa(age)
	fmt.Println("My age is "+age_str)
	//or
	fmt.Println("My age is "+strconv.Itoa(age))
	a := "23"
	age_int, _ := strconv.Atoi(a) //Atoi returns 2 values, in this case, _ dismisses 2nd one
	fmt.Println( age_int + 10 )
}