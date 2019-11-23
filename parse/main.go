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
	s := bufio.NewScanner(r)

	for s.Scan() {
		fmt.Printf("read: %s ", s.Text())
	}
	fmt.Println("----------------------------- ")
}
