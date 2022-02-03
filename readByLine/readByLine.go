package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func lineByLine(file string) error {
	var err error
	fd, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fd.Close()

	reader := bufio.NewReader(fd)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("Error reading file %s", err)
			break
		}
		fmt.Print(line)
	}
	return nil
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Printf("usage: lByL <file1> [<file2> ...]\n")
		return
	}

	for _, file := range flag.Args() {
		err := lineByLine(file)
		if err != nil {
			fmt.Println(err)
		}
	}
}
