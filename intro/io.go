package main

import (
	"io"
	"strings"
	"bytes"
)

func WriteString(s string, w io.Writer) error {
	if _, err := w.Write([]byte(s)); err != nil {
		return err
	}
	return nil
}

func ReadString(r io.Reader) (string, error) {
	data := make([]byte, 2048)
	bytesRead, err := r.Read(data)
	if err != nil && err != io.EOF {
		return "", err
	}
	return string(data[:bytesRead]), nil
}

type UpperWriter struct {
   UpperString string
}

func (u *UpperWriter) Write(p []byte) (int, error) {
   u.UpperString = strings.ToUpper(string(p))
   return len(p), nil
}

func Copy(r io.Reader, w io.Writer, n uint) (error) {
	data, err := ReadString(r)
	if err != nil {
		return err
	}
	if n < uint(len(data)) {
		data = data[:n]
	}
	return WriteString(data, w)
}

func Contains(r io.Reader, seq []byte) (bool, error) {
    buf := make([]byte, 2 * len(seq))
	n, err := r.Read(buf)
	if err != nil {
		return false, err
	}
	return bytes.Contains(buf[:n], seq), nil
}