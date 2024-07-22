package main

import "fmt"

func main()  {
	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	// basic for loop
	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	// for range loop
	for i := range colors {
		fmt.Println(colors[i])
	}

	// for each loop
	for _, color := range colors {
		fmt.Println(color)
	}

	// equivalent while loop
	value := 1
	for value < 10 {
		fmt.Println("Value:", value)
		value++
	}
	// Value: 1
	// Value: 2
	// Value: 3
	// Value: 4
	// Value: 5
	// Value: 6
	// Value: 7
	// Value: 8
	// Value: 9

	// Go labels, it replaces break and continue
	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println("Sum:", sum)
		if sum > 200 {
			goto theEnd
		}
	}

	theEnd:
		fmt.Println("End of the program")
	// Sum: 2
	// Sum: 4
	// Sum: 8
	// Sum: 16
	// Sum: 32
	// Sum: 64
	// Sum: 128
	// Sum: 256
	// Sum: 512
	// Sum: 1024
	// End of the program

	// Go supports break and continue statements
	product := 2
	for product < 100000 {
		product *= product
		fmt.Println("Product:", product)
		if product > 30 {
			fmt.Println("Break statement")
			break
		}
	}
	// Product: 4
	// Product: 16
	// Product: 256
	// Break statement
}