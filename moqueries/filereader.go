package moqueries

import (
	"io"
	"os"

	moqos "moqueries.org/std/os"
)

type OpenFn = moqos.Open_genType

type FileToReaderFn func(file *os.File) io.Reader

var FileToReader = func(file *os.File) io.Reader {
	return file
}

type FileReader struct {
	openFn   OpenFn
	readerFn FileToReaderFn
}

func New(openFn OpenFn, readerFn FileToReaderFn) *FileReader {
	return &FileReader{
		openFn:   openFn,
		readerFn: readerFn,
	}
}

func (r *FileReader) Read(name string) (string, error) {
	f, err := r.openFn(name)
	if err != nil {
		return "", err
	}

	contents, err := io.ReadAll(r.readerFn(f))
	if err != nil {
		return "", err
	}

	return string(contents), nil
}
