package main

import "fmt"

func main() {
	s := "Hello World"
	tricky_string := "🧁 Yüm! 😋"
	printLetters(s)
	printLetters(tricky_string)
	twoSum([]int{2, 7, 11, 15}, 9)
}

func printLetters(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	fmt.Println("\n----\n")

	for i := 0; i < len(s); i++ {
		fmt.Println(string(s[i]))
	}

	fmt.Println("\n----\n")

	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}

	fmt.Println("\n----\n")

	asRunes := []rune(s)

	fmt.Println("\n----\n")

	for i := 0; i < len(asRunes); i++ {
		fmt.Printf("%c ", asRunes[i])
	}

	fmt.Println("\n----\n")

	for _, letter := range s {
		fmt.Println(string(letter))
	}

}

func twoSum(nums []int, target int) []int {
	myMap := make(map[int]int)
	for i, num := range nums {
		goal := target - num
		value, present := myMap[goal]
		if present {
			return []int{value, i}
		}
		myMap[num] = i
	}
	return []int{-1, -1}
}
