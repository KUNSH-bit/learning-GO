package main

import "fmt"

func concat(s1 string, s2 string)string{ // called function signature
	return s1+s2;
}

func add(x,y int)int{ // if both parameters are of same type, we can write like this
	return x+y;
}
func increment(x int){ // pass by value, so x will not be changed
	x++;  // no return so x in main will remain unchanged
}

func properIncrement(x int)int{ 
	x++ ;
	return x;
}
//go doesnt allow u to keep unused variables at all
func name(string)(string,string){ // this way we can return multiple values from a function
	return "Kunsh","Agarwal"
}

func main() {

	fmt.Println(concat("Hello ","World"));
	fmt.Println(add(10,20));
	x:=10;
	increment(x); 
	fmt.Println(x);
	y := properIncrement(10);
	fmt.Println(y); 
	firstName, _ := name("Kunsh"); // _ is used to ignore the second return value
	fmt.Println(firstName);
}