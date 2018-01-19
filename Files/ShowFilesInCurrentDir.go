package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/**
 * Created by Pan on 1/9/2018.
 */
func main() {
	dir, err := os.Open(".")
	if err != nil {
		fmt.Println("Error in opening the dir file")
	}
	defer dir.close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("Error in getting file infos inside the dir file")
	}

	for _, fileInfo := range fileInfos {
		fileName := fileInfo.Name()
		if strings.HasSuffix(fileName, ".go") {
			//Prints the content of GoLang file which contains in the current dir.
			//Will also display this file (ShowFilesInCurrentDir.go) content
			bytes, err := ioutil.ReadFile(fileName)
			if err != nil {
				fmt.Println("Error in converting file content to bytes")
			}

			fmt.Println(fileName)
			fmt.Println(string(bytes))
		}
	}
}
