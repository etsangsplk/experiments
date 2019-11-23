package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Reading expression input: > ")
	fmt.Println("----------------------------- ")
	r := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		fmt.Printf("read: %s ", scanner.Text())
	}

	fmt.Println("-------Done processing--------- ")

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "reading standard input: ", err)
	}
}
