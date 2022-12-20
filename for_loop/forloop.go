package main

import (
	"fmt"
)

func main() {

	// for loop
	for i := 0; i < 10; i++ {
		fmt.Printf("Simple loop %v \n", i)
	}
	z := 1
	for z <= 5 {
		if z == 5 {
			break
		}
		fmt.Printf("z value %v \n", z)
		z += 1
	}

	//grading indentifier

	gradeArray := []int{40, 50, 60, 80, 90}

	for i := 0; i < len(gradeArray); i++ {
		if gradeArray[i] < 45 {
			fmt.Println("You  Scored C grade")
		} else if gradeArray[i] < 55 {
			fmt.Println("You  Scored B grade")
		} else if gradeArray[i] < 65 {
			fmt.Println("You  Scored A- grade")
		} else if gradeArray[i] < 85 {
			fmt.Println("You Scored A grade")
		} else {
			fmt.Println("You scored A+ grade")
		}

	}

}
