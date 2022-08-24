package main

import "fmt"

func main() {

	score := 94
	highScore := 99
	// println will always print the newline
	// at the end of statement automatically
	fmt.Println("My score is", score)

	// use format spcifiers
	// use printf with format specifiers. Cannot use Println

	// %v is one of many format specifiers!

	// printf will always print the newline
	// at the end of statement automatically
	// use \n to print the new line
	fmt.Printf("My score is %v. \n", score)
	fmt.Printf("My score is %v \n", score)
	fmt.Printf("My score is '%v' \n", score)
	fmt.Printf("My score: %v. \n", score)

	// use multiple formst specifiers in the single statement
	fmt.Printf("My score is %v. My highest score till date is %v", score, score)
	fmt.Printf("My score is %v. Highest score till date is %v \n", score, highScore)

	// %T -> datatype
	fmt.Printf("Value stores in score variable is %v. Data type of score is %T \n", score, score)
	// working with space in format specifiers
	fmt.Printf("score is %v.\n", score)    //right align
	fmt.Printf("score is %-10v.\n", score) // left align
}
