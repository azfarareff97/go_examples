package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		fmt.Printf("20 x %v = %v \n", i, i*20)
	}
	fmt.Println("-------------------------------")
	var scores = [8]int{4, 6, 8, 2, 0, 3, 5, 8}
	//iterate through the scores array
	for index, value := range scores {
		fmt.Printf("Index: %v, value: %v \n", index, value)
	}

}
