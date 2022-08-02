package main

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

type filebuffer struct {
	*os.File
	deleted bool
}

func newFilebuffer(f *os.File) *filebuffer {
	return &filebuffer{f, false}
}

var _ io.ReadSeekCloser = &filebuffer{}

func (fbuff *filebuffer) Close() (err error) {
	err = fbuff.File.Close()
	if fbuff.deleted {
		return
	}

	if errm := os.Remove(fbuff.File.Name()); errm != nil {
		log.Debugf("Error removing temporary file %s: %v", fbuff.File.Name(), errm)
		return
	}

	log.Debugf("Removed temporary file %s", fbuff.File.Name())
	fbuff.deleted = true
	return
}
