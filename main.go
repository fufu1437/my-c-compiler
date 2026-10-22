package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// TODO (compiler-pipeline): implement per the lesson description.

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			continue
		}

		files := strings.Split(line, ".")
		// if == "c" {
		// }
		switch files[len(files)-1] {
		case "c":
			fmt.Println("SOURCE")
		case "i":
			fmt.Println("PREPROCESSED")
		case "s":
			fmt.Println("ASSEMBLY")
		case "o":
			fmt.Println("OBJECT")
		case "a":
			fmt.Println("STATIC_LIB")
		case "so", "dylib", "dll":
			fmt.Println("SHARED_LIB")
		default:
			fmt.Println("EXECUTABLE")
		}
	}
}
