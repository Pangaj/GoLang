package main

import "fmt"

/**
 * Created by Pan on 1/13/2018.
 */
func main() {
	// A slice is a segment of an array.
	// Like arrays slices are indexable and have a length.
	// Unlike arrays this length is allowed to change

	//Examples on how to initialise Slices
	initVariable := make([]float32,5)
	fmt.Println("value : ", initVariable)
	fmt.Println("length : ", len(initVariable))
	fmt.Println("Capacity : ", cap(initVariable))

	capVariable := make([]float32,5,10)
	fmt.Println("value : ", capVariable)
	fmt.Println("length : ", len(capVariable))
	fmt.Println("Capacity : ", cap(capVariable))
}
