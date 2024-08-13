package main

import "fmt"

func romanToInt(s string) int {
	value_map := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	total_sum := 0
	last_value := 0

	for _, elem := range s {
		current_value := value_map[string(elem)]

		if current_value > last_value{
			total_sum = total_sum - last_value + current_value - last_value
		} else if current_value <= last_value {
			total_sum += current_value

		}

		last_value = current_value 
	}

	return total_sum
}

func main() {
	value := romanToInt("MCMXCIV")
	fmt.Println(value)
}
