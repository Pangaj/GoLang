package main

import "fmt"

/**
 * Created by Pan on 1/11/2018.
 */
func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, " - Even")
		} else {
			fmt.Println(i, " - Odd")
		}
	}
}
