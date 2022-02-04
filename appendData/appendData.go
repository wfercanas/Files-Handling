package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args
	if len(arguments) != 3 {
		fmt.Printf("usage: %s message filename\n", filepath.Base(arguments[0]))
		return
	}

	message := arguments[1]
	filename := arguments[2]

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(file, "%s\n", message)
}
