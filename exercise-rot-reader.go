package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

// ROT13変換関数
func rot13(b byte) byte {
	switch {
	case 'a' <= b && b <= 'z':
		// 'a'から'z'の場合、13文字ずらす
		return 'a' + (b-'a'+13)%26
	case 'A' <= b && b <= 'Z':
		// 'A'から'Z'の場合、13文字ずらす
		return 'A' + (b-'A'+13)%26
	default:
		// アルファベット以外はそのまま
		return b
	}
}

// Readメソッドの実装
func (rot13r *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot13r.r.Read(p)
	if err != nil {
		return n, err
	}

	for i := range p[:n] {
		p[i] = rot13(p[i])
	}

	return n, nil
}
