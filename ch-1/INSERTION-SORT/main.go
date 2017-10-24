package main

import (
	"fmt"
	"time"
	"math/rand"
)

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func Sort(in []int) []int {
	if in == nil {
		return nil
	}
	for i := 1; i < len(in); i++ {
		current := in[i]
		j := i - 1
		for j >= 0 && in[j] > current {
			in[j+1] = in[j]
			j = j - 1
		}
		in[j+1] = current
	}
	return in
}


func main() {
	slice := generateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	Sort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")

}
