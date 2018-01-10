package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

/**
 * Created by Pan on 1/10/2018.
 */
func main() {
	file, err := os.OpenFile("NewFile.txt", os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("\nFile error while opening the file : ", err)
	}

	_, err = file.WriteString("\nBefore getting file info. Don't append any content after getting the file stats")
	if err != nil {
		log.Fatal(err)
	}

	/**
	 * Write or append any content before getting the file.Stat() function :)
	 */
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("\nError in receving FileInfo describing file : ", err)
	}

	openFileBytes := make([]byte, fileInfo.Size())

	//If you write some content, the file offset will move to end of the file.
	//So inOrder to move to first we use file.Seek(... , ...)
	file.Seek(0, io.SeekStart)

	_, err = file.Read(openFileBytes)
	if err != nil {
		fmt.Println("Error in reading file : ", err)
	}

	fmt.Println(string(openFileBytes))
}
