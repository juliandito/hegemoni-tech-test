package main

import (
	"errors"
	"fmt"
)

func CalculateAverage(numbers []int) (float64, error) {
	if len(numbers) == 0 {
		return 0.0, errors.New("cannot calculate average of an empty slice")
	}

	total := 0
	for _, number := range numbers {
		total += number
	}

	average := float64(total) / float64(len(numbers))
	return average, nil
}

func main() {
	sample := []int{10, 20, 30, 40}
	average, err := CalculateAverage(sample)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Average:", average)
	}

	empty := []int{}
	_, err = CalculateAverage(empty) // print only err
	if err != nil {
		fmt.Println("Error:", err)
	}
}
