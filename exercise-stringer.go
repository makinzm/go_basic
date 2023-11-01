package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

// https://go-tour-jp.appspot.com/methods/18

func (ip IPAddr) String() string {
	var res string
	for i, v := range ip{
		if i < 3 {
			res += (strconv.Itoa(int(v)) + ".")
		} else {
			res += strconv.Itoa(int(v))
		}
	}
	return res
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
