package basics

import "fmt"

func main() {

	// Simple for loop
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// iterate over collection
	// Printf for specific format, %v for general value, %d for number
	numbers := []int{1, 2, 3, 4, 5, 6}
	for index, value := range numbers {
		// Skip this specific iteration
		if index == 2 {
			continue
		}

		// Stop at this point (iteration)
		if index == 4 {
			break
		}

		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	for i := 0; i <= 5; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println("Is even: ", i)
	}

	rows := 5

	// Outer loop
	for i := 1; i <= rows; i++ {
		// Inner loops
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// Loop within range instead of i := 1; i < 10; i++
	for i := range 10 {
		fmt.Println(i)
	}
}
