package main

import (
	"fmt"
)

func main() {

	/*var operator string
	var num1, num2 int

	fmt.Print("Enter fisrt numbers: ")
	fmt.Scan(&num1)

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scan(&operator)

	fmt.Print("Enter second numbers: ")
	fmt.Scan(&num2)

	switch operator {
	case "+":
		fmt.Printf("%d + %d = %d\n", num1, num2, num1+num2)
	case "-":
		fmt.Printf("%d - %d = %d\n", num1, num2, num1-num2)
	case "*":
		fmt.Printf("%d * %d = %d\n", num1, num2, num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Printf("%d / %d = %d\n", num1, num2, num1/num2)
		} else {
			fmt.Println("Error: Division by zero")
		}
	default:
		fmt.Println("Error: Invalid operator")
	} */

	/*var gcpProjects = [3]string{"project1", "project2", "project3"}
	fmt.Println(gcpProjects[0:1])
	for _, project := range gcpProjects {
		fmt.Println(project)
	}
	awsProjects := [3]string{"aws-project1", "aws-project2", "aws-project3"}
	fmt.Println(awsProjects[0:1])

	var numbers []int
	name := []string{"John", "Doe", "Smith"}
	fmt.Println(name)
	fmt.Println(numbers)
	fmt.Println(len(name))
	fmt.Println(cap(name))

	numbers = append(numbers, 1)
	numbers = append(numbers, 2)
	numbers = append(numbers, 3)
	numbers = append(numbers, 4)
	numbers = append(numbers, 5)

	name = append(name, "Jane")

	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))

	numbers = append(numbers, 6)
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))

	length := len(name)
	fmt.Println(length)

	var ages = map[string]int{
		"John":  25,
		"Doe":   30,
		"Smith": 35,
	}

	fmt.Println(ages)
	fmt.Println(ages["John"])

	fmt.Println(len(ages))

	ages["Jane"] = 28
	fmt.Println(ages)
	fmt.Println(len(ages))
	delete(ages, "Doe")
	fmt.Println(ages)
	fmt.Println(len(ages))

	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}*/

	// Lopps

	rows := 9
	for i := 1; i <= rows; i++ {
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			if i == 9 && j == 4 {
				continue
			}
			fmt.Print("*")
		}
		fmt.Println()
	}

	rowsInsvers := 9
	for i := rowsInsvers; i >= 1; i-- {
		//space
		for j := 1; j <= rowsInsvers-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			if i == 8 && (k <= 4 && k >= 8) {
				fmt.Print(" ")
				break
			}
			fmt.Print("*")
		}
		fmt.Println()
	}

}
