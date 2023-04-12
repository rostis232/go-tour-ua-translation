// +build no-build OMIT

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

<<<<<<< HEAD
// TODO: Створіть метод Read([]byte) (int, error) для структурного типу MyReader.
=======
// TODO: Створіть функцію-метод Read([]byte) (int, error) для структурного типу MyReader.
>>>>>>> c00883de89fe4ae0840173dc8f913ce8f199dc81

func main() {
	reader.Validate(MyReader{})
}
