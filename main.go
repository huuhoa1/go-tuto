package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	/* f, err := os.Open("letters.txt")
	if err != nil {
		panic(err)
	} */
	r := strings.NewReader("Hello worldddd")
	n, err := countAlphabets(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Letters: %d\n", n)
}

func countAlphabets(r io.Reader) (int, error) {
	count := 0
	buffer := make([]byte, 1024)

	for {
		n, err := r.Read(buffer)       // infinite loop, read 1024 chars at a time
		for _, l := range buffer[:n] { //loop on the n chars read
			if (l >= 'A' && l <= 'Z') || (l >= 'a' && l <= 'z') {
				count++
			}
		}
		if err == io.EOF {
			return count, nil
		}
		if err != nil {
			return 0, err
		}
	}
}
