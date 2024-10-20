package printer

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func PrintHello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func Length(a int) string {
	switch {
	case a < 0:
		return "negative"
	case a == 0:
		return "zero"
	case a < 10:
		return "short"
	case a < 100:
		return "long"
	}
	return "very long"
}

var ErrInvalidUTF8 = errors.New("invalid utf8")

func GetUTFLength(input []byte) (int, error) {
	if !utf8.Valid(input) {
		return 0, ErrInvalidUTF8
	}

	return utf8.RuneCount(input), nil
}

func main() {
	t := []byte("Hello, World!")
	fmt.Println(GetUTFLength(t))
}
