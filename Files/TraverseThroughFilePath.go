package main

import (
	"fmt"
	"os"
	"path/filepath"
)

/**
 * Created by Pan on 1/9/2018.
 */
func main() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(info.Name())
		return nil
	})
}
