package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("result.txt") // create or truncate file
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("sample.txt") // write new content
	if err != nil {
		fmt.Println("Failed to write to file:", err)
		return
	}

	fmt.Println("Success!")
}
