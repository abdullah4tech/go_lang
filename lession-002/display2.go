package main
import "fmt"

/*
  Author: Abdullah Oluwatomisin Mustapha
  Date: March 27, 2026


  Description
  	Exploring different ways to display and format string using the Print and Println commands

 */

func main() {
	fmt.Println("Hello" + "world") // joins the two strings together
	fmt.Println("Hello", "world") // join the two string together and give a space in between them

	fmt.Println(1+1)

	fmt.Println(3.0/7.0)

	fmt.Println(20/5)

	fmt.Println(true && true)

	fmt.Println(true && false)

	fmt.Println(false && true)

	fmt.Println(true || false)

	fmt.Println(false || true)
}
