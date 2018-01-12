package main

import "fmt"

/**
 * Created by Pan on 1/11/2018.
 */
func main() {
	//An array is a numbered sequence of elements of a single type with a fixed length
	var arr [5]int

	arr[0] = 1
	arr[1] = 2
	arr[3] = 4
	arr[4] = 5

	/*
	 *	Or directly initialise as below
	 * 	var arr = [5]int{1,2,3,4,5}
	 *  arr := [5]int{1,2,3,4,5}
	 * 	arr := [5]int{
	 *				1,
	 *				2,
	 *				3,
	 *				4,
	 *				5,
	 *     		}
	 */

	//output : [1 2 0 4 5]
	fmt.Println(arr)

	//output : 5
	fmt.Println(arr[4])
}
