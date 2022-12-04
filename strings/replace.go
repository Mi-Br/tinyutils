package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "not enough arguments, expect 2: old and new keywords for replacement")
		os.Exit(-1)
	}

	old, new := os.Args[1], os.Args[2]
	scan := bufio.NewScanner(os.Stdin) //scanner on terminal input

	for scan.Scan() {
		s := strings.Split(scan.Text(), old)
		out := strings.Join(s, new)
		fmt.Println(out)
	}

}
