package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

/**
 * Created by Pan on 1/8/2018.
 */
func main() {
	/**
	 *There are 2 ways to open and write content in file
	 *They are,
	 *1. using Open and Write command. But it is hectic to do this often. So there is another simple way to do so.
	 *2. using ReadFile command.
	 *In this program we will show you the both types.
	 */

	/*Type One - Using Open command*/
	fmt.Println("Type 1 :- Using Open Command")

	file, err := os.Open("NewFile.txt")
	if err != nil {
		fmt.Println("File error while opening the file : ", err)
	}

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error in receving FileInfo describing file : ", err)
	}

	openFileBytes := make([]byte, fileInfo.Size())

	_, err = file.Read(openFileBytes)
	if err != nil {
		fmt.Println("Error in reading file : ", err)
	}

	fmt.Println(string(openFileBytes))

	/*Type Two - Using ReadFile command*/
	fmt.Println("Type 2 :- Using ReadFile command")

	readFileBytes, err := ioutil.ReadFile("NewFile.txt")
	if err != nil {
		fmt.Println("Error in reading file :", err)
	}

	fmt.Println(string(readFileBytes))
}
