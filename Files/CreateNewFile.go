package main

import (
	"fmt"
	"os"
)

/**
 * Created by Pan on 1/8/2018.
 */
func main() {
	//os.Create will create a new file
	//If already exists, then truncates the file
	file, err := os.Create("NewFile.txt")

	if err != nil {
		fmt.Println("Error in creating file")
	}
	defer file.close()

	_,err = file.Write([]byte("New File Created :)"))
	if err != nil {
		fmt.Println("Error in writing into the file :(")
	}
}
