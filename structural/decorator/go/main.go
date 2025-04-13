package main

import "fmt"

func main() {
	pizza := &VeggieMania{}

	// Add cheese
	pizzaWithCheese := &CheeseTopping{pizza}

	// Add tomato
	pizzaWithCheeseAndTomato := &TomatoTopping{pizzaWithCheese}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
}
