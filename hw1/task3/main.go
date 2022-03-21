package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	files_count := 1000000
	folder_name := "./temp"
	err := os.Mkdir(folder_name, 0755)
	if err != nil {
		fmt.Print(err)
	}

	for i := 1; i <= files_count; i++ {
		file_name := folder_name + "/" + "emptyFile" + strconv.Itoa(i) + ".txt"
		emptyFile, err := os.Create(file_name)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer emptyFile.Close()
	}

}
