package printer

import (
	"testing"
)

func TestPrintHello(t *testing.T) {
	got := PrintHello("Igor")
	expected := "Hello, Igor!"

	if got != expected {
		t.Fatalf(`PrintHello("Igor") = %q, want %q`, got, expected)
	}
}

func TestGetUTFLength(t *testing.T) {
	cases := []struct {
		in  []byte
		out int
		err error
	}{
		{[]byte("Привет, мир!"), 12, nil},
		{[]byte{0xff, 0xfe, 0xfd}, 0, ErrInvalidUTF8},
		{[]byte(""), 0, nil},
		{[]byte("Go語"), 3, nil},
	}

	for _, tc := range cases {
		t.Run(string(tc.in), func(t *testing.T) {
			got, err := GetUTFLength(tc.in)
			if got != tc.out || err != tc.err {
				t.Errorf("GetUTFLength(%q) = %d, %v; want %d, %v", tc.in, got, err, tc.out, tc.err)
			}
		})
	}
}
