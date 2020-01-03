package main

import (
	"math/rand"
	"time"
	"fmt"
)

func random_sequence(minimum, maximum int) []int { //returns a shuffled array
	array := make([]int, 0)
	for i := minimum; i < maximum; i++ { //create array and append numbers to it
		array = append(array, i)
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(array), func(i, j int) { array[i], array[j] = array[j], array[i] }) //shuffle the array

	return array
}

func max(array []int) int {
	largest := array[0]
	for _, i := range array {
		if i > largest {
			largest = i
		}
	}
	return largest
}

func counting_sort(array []int) []int {
	count := make([]int, max(array)+1)

	for i := 0; i < len(array); i++ {
		count[array[i]]++
	}

	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	output := make([]int, len(array))
	for _, i := range array {
		output[count[i]-1] = i
        count[i]--
	}

	return output
}

func main() {
	shuffled_array := random_sequence(0, 1000)
	fmt.Println("COUNTING SORT")
	fmt.Println(shuffled_array)
	fmt.Println()
	fmt.Println(counting_sort(shuffled_array))
}