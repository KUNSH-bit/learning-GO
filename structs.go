// How to use structs, switch, and nested/embedded structs in Go
package main

import "fmt"

type Classification struct {
	Archetype string
}

type Car struct {
	Brand          string
	Model          string
	Year           int
	Classification Classification
}

// instead u can also write it like this tho it is not recommended to use nested structs like this
// type Car struct {
// 		Brand string
// 		Model string
// 		Year int
// 		Classification struct{
// 			Archetype string
// 		}
// 	}

type embed struct {
	Car
	something string
}

func printClassification(c Car) {
	switch c.Classification.Archetype {
	case "classic":
		fmt.Printf("%s %s is a Classic Sports Car\n", c.Brand, c.Model)
	case "hypercar":
		fmt.Printf("%s %s is a Hypercar\n", c.Brand, c.Model)
	case "JDM":
		fmt.Printf("%s %s is a JDM Car\n", c.Brand, c.Model)
	default:
		fmt.Printf("%s %s has an unknown archetype (%q)\n", c.Brand, c.Model, c.Classification.Archetype)
	}
}

func describeCar(c Car) {
	fmt.Printf("%s made a well known car, the %s %s, in %d\n", c.Brand, c.Brand, c.Model, c.Year)
	printClassification(c)
}

func main() {
	e := embed{
		Car: Car{
			Brand:          "Porsche",
			Model:          "912",
			Year:           1965,
			Classification: Classification{Archetype: "classic"},
		},
		something: "extra detail",
	}

	cars := []Car{
		{Brand: "Porsche", Model: "912", Year: 1965, Classification: Classification{Archetype: "classic"}},
		{Brand: "Aston Martin", Model: "Valhalla", Year: 2022, Classification: Classification{Archetype: "hypercar"}},
		{Brand: "Mazda", Model: "RX-7", Year: 1985, Classification: Classification{Archetype: "JDM"}},
	}

	fmt.Println(e.Car.Brand) // no need to use e.Car.Brand since the Car struct is embedded in the embed struct
	fmt.Println(e.Brand)     
	fmt.Println(e.something) 
	describeCar(e.Car)

	for _, c := range cars {
		describeCar(c)
	}
}