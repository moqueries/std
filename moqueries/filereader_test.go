package moqueries_test

import (
	"errors"
	"io"
	"os"
	"testing"

	"moqueries.org/runtime/moq"

	moqio "moqueries.org/std/io"
	"moqueries.org/std/moqueries"
	moqos "moqueries.org/std/os"
)

func TestRead(t *testing.T) {
	var (
		scene      *moq.Scene
		mockOpenFn *moqos.MoqOpen_genType
		mockReader *moqio.MoqReader

		file *os.File

		reader *moqueries.FileReader
	)

	beforeEach := func(t *testing.T) {
		scene = moq.NewScene(t)
		mockOpenFn = moqos.NewMoqOpen_genType(scene, nil)
		mockReader = moqio.NewMoqReader(scene, nil)

		file = &os.File{}
		reader = moqueries.New(mockOpenFn.Mock(), func(f *os.File) io.Reader {
			if f != file {
				t.Fatalf("got %#v, want %#v", f, file)
			}
			return mockReader.Mock()
		})
	}

	t.Run("happy path", func(t *testing.T) {
		// ASSEMBLE
		beforeEach(t)
		defer scene.AssertExpectationsMet()

		mockOpenFn.OnCall("/this/file").ReturnResults(file, nil)
		mockReader.OnCall().Read(nil).Any().P().
			DoReturnResults(func(p []byte) (n int, err error) {
				out := []byte("Hello")
				copy(p, out)
				return len(out), nil
			})
		mockReader.OnCall().Read(nil).Any().P().ReturnResults(0, io.EOF)

		// ACT
		contents, err := reader.Read("/this/file")

		// ASSERT
		if err != nil {
			t.Fatalf("got %#v, want no error", err)
		}
		if contents != "Hello" {
			t.Errorf("got %s, want Hello", contents)
		}
	})

	t.Run("open error", func(t *testing.T) {
		// ASSEMBLE
		beforeEach(t)
		defer scene.AssertExpectationsMet()

		mockOpenFn.OnCall("/this/file").ReturnResults(nil, errors.New("open error"))

		// ACT
		contents, err := reader.Read("/this/file")

		// ASSERT
		if err == nil {
			t.Fatalf("got no error, want error")
		}
		if err.Error() != "open error" {
			t.Errorf("got %s, want 'open error'", err.Error())
		}
		if contents != "" {
			t.Errorf("got %s, want ''", contents)
		}
	})

	t.Run("read error", func(t *testing.T) {
		// ASSEMBLE
		beforeEach(t)
		defer scene.AssertExpectationsMet()

		mockOpenFn.OnCall("/this/file").ReturnResults(file, nil)
		mockReader.OnCall().Read(nil).Any().P().ReturnResults(0, errors.New("read error"))

		// ACT
		contents, err := reader.Read("/this/file")

		// ASSERT
		if err == nil {
			t.Fatalf("got no error, want error")
		}
		if err.Error() != "read error" {
			t.Errorf("got %s, want 'read error'", err.Error())
		}
		if contents != "" {
			t.Errorf("got %s, want ''", contents)
		}
	})
}
