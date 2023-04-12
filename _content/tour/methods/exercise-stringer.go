// +build OMIT

package main

import "fmt"

type IPAddr [4]byte

<<<<<<< HEAD
// TODO: Створіть метод "String() string" для структурного типу IPAddr.
=======
// TODO: Створіть функцію-метод "String() string" для структурного типу IPAddr.
>>>>>>> c00883de89fe4ae0840173dc8f913ce8f199dc81

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
