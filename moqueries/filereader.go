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

	// TODO: Use io.ReadAll when older standard libraries are behind us
	// b, err := io.ReadAll(r.readerFn(f))
	// if err != nil {
	// 	return "", err
	// }
	// return string(b), nil
	b := make([]byte, 0, 512)
	for {
		if len(b) == cap(b) {
			// Add more capacity (let append pick how much).
			b = append(b, 0)[:len(b)]
		}
		n, err := r.readerFn(f).Read(b[len(b):cap(b)])
		b = b[:len(b)+n]
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return string(b), err
		}
	}
}
