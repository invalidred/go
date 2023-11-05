package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "errorL", err)
		os.Exit(1)
	}
}

/*
➜  go_by_examples git:(main) ✗ cat line-filter.txt | go run line-filters.go
HELLO
FILTER
*/
