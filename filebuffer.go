package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
)

var _ io.ReadSeekCloser = &filebuffer{}

type filebuffer struct {
	*os.File
	deleted bool
}

func newFilebuffer() (*filebuffer, error) {
	f, err := ioutil.TempFile(os.TempDir(), "s3-proxy-*")
	if err != nil {
		return nil, fmt.Errorf("unable to create temporary file: %w", err)
	}
	return &filebuffer{f, false}, nil
}

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
