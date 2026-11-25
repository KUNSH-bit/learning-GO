// same as any other language just without ().
package main

import "fmt"

func main() {
	const age=21;

	if age>=18 {
		fmt.Println("You are eligible to vote");
	} else if age>10000{
		fmt.Println("How??");
	}else{
		fmt.Println("You are not eligible to vote");
	}

}