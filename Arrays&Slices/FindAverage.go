package main

import "fmt"

/**
 * Created by Pan on 1/11/2018.
 */
func main() {
	var arr = [5]int{5,0,0,0,0}
	//arr := [5]int{1,2,3,4,5}

	var total int
	for i:= 0 ; i<len(arr);i++ {
		total += arr[i]
	}

	average := float32(total)/float32(len(arr))
	fmt.Println(average)
}
