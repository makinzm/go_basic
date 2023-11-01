package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// https://go-tour-jp.appspot.com/methods/22
func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A' // We cannot use "A" because of byte
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
