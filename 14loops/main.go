package main

import "fmt"

func main() {
	fmt.Println("Loops in Go")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i, d := range days {
	// 	fmt.Printf("Index: %v, Day: %v\n", i, d)
	// }

	rogueValue := 1

	for rogueValue < 10 {
		if rogueValue == 5 {
			rogueValue++
			continue
		}

		if rogueValue == 8 {
			goto lco
		}

		fmt.Println(rogueValue)
		rogueValue++
	}

lco:
	fmt.Println("I am outside the loop")
}
