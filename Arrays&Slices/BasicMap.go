package main

import "fmt"

/**
 * Created by Pan on 1/14/2018.
 */
func main() {
	basicMap := make(map[string]int)

	basicMap["one"] = 1
	basicMap["two"] = 2
	basicMap["three"] = 1
	fmt.Println(basicMap["one"])
	delete(basicMap, "one")
	fmt.Println(basicMap["one"])
}
