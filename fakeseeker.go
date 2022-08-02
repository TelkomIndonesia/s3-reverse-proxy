package main

import "io"

type fakeseeker struct {
	io.Reader
}

func (fakeseeker) Seek(offset int64, whence int) (int64, error) {
	panic("should not be called")
}
