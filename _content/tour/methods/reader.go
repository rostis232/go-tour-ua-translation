//go:build OMIT
// +build OMIT

package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
<<<<<<< HEAD
	r := strings.NewReader("Вітаю, Reader!")

=======
	r := strings.NewReader("Привіт, Reader!")
>>>>>>> c00883de89fe4ae0840173dc8f913ce8f199dc81
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
