package main

import "fmt"

/**
 * Created by Pan on 1/12/2018.
 */
func main() {
	var arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, index := range arr {
		fmt.Println(index)
	}
}
