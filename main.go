package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	/* f, err := os.Open("letters.txt")
	if err != nil {
		panic(err)
	} */
	/* r := strings.NewReader("Hello worldddd")
	n, err := countAlphabets(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Letters: %d\n", n) */
	f, err := os.Create("writing.txt")
	if err != nil {
		panic(err)
	}
	defer (*f).Close()

	n, err := writeString("Hello Can Huynh!", f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("written bytes: %d\n", n)

}

func writeString(s string, w io.Writer) (int, error) {
	n, err := w.Write([]byte(s))
	if err != nil {
		return 0, fmt.Errorf("error occurred while writing: %w", err)
	}
	return n, nil
}
