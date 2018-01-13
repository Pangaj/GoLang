package main

import "fmt"

/**
 * Created by Pan on 1/13/2018.
 */
func main() {
	reSliceArray := make([]int, 0, 10)

	for i:= 0; i < cap(reSliceArray);i++ {
		reSliceArray = reSliceArray[0 : i+1]
		reSliceArray[i] = i+1
	}

	for _, index := range reSliceArray {
		fmt.Println(index)
	}
}