package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Printf("I am %v.","Kunsh")
	fmt.Printf("\nI am %v years old\n",20)
	fmt.Printf("%%f prints float = %f \nwhile, %%.2f prints float with 2 decimal places = %.2f",3.14159,3.14159)

	const data="data type";
	msg := fmt.Sprintf("automatically takes %s, so no need to declare", data);
	fmt.Printf("\n%s\n", msg);
};