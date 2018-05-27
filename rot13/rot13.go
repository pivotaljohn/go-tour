package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(buff []byte) (int, error) {
	for i, b := range byte {
		if b >= rune('A') && b <= rune('Z') {
			fmt.Println("U")
		} else if b >= rune('a') && b <= rune('z') {
			fmt.Println("l")
		} else {
			fmt.Println(".")
		}
	}
	return len(buff), nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
