package main

import "fmt"

func main() {
	// const full_name string = "Golang"
	// fmt.Printf("Hello %s", full_name)

	// var value = 5 + 3*2
	// fmt.Println(value)

	var firstCondition bool = 3 < 5
	var secondCondition bool = "golang" == "Golang"
	var thirdCondition bool = 10 != 3.8
	var fourthCondition bool = 11 <= 11

	fmt.Println("firstCondition ==>", firstCondition)
	fmt.Println("secondCondition ==>", secondCondition)
	fmt.Println("thirdCondition ==>", thirdCondition)
	fmt.Println("fourthCondition ==>", fourthCondition)

	var right = true
	var left = false

	var leftAndRight = right && left
	fmt.Printf("left && right \t(%t)\n", leftAndRight)

	var leftOrRight = right || left
	fmt.Printf("left || right \t(%t)\n", leftOrRight)

	var notRight = !right
	fmt.Printf("!right \t(%t)\n", notRight)
}
