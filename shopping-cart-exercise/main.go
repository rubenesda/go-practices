package main

import "fmt"

func main()  {
	var cart []CartItem
	var apples = CartItem{"apple", 2.99, 3}
	var oranges = CartItem{"orange", 1.50, 8}
	var bananas = CartItem{"banana", .49, 12}
	cart = append(cart, apples, oranges, bananas)
	result := calculateTotal(cart)
	fmt.Printf("Cart Total: %.6f\n", result)
}

func calculateTotal(cart []CartItem) float64 {
	var sum float64 = 0
    for _, item := range cart {
        sum += item.price * float64(item.quantity)
    }
  return sum
}

type CartItem struct {
	name string
	price float64
	quantity int
}