package httputil2

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Returns a new http.FileSystem that is first loaded into memory.
//
// Great when Disk IO is very slow.
func MemDir(path string) http.FileSystem {
	fs := NewMemoryFileSystem()
	err := fs.AddDir(path)
	if err != nil {
		log.Fatal("Fatal", err)
	}

	return fs
}

type MemoryFileSystem map[string]*MemoryFile

func NewMemoryFileSystem() MemoryFileSystem {
	return make(MemoryFileSystem)
}

func (self MemoryFileSystem) Open(name string) (http.File, error) {
	name = filepath.Clean(name)

	x := self[name]
	if x != nil {
		if x.err != nil {
			return nil, x.err
		}

		return &MemRef{bytes.NewReader(x.data), x, 0}, nil
	}
	return nil, &os.PathError{
		Op:   "open",
		Path: name,
		Err:  fmt.Errorf("no such file or directory"),
	}
}

func (self MemoryFileSystem) AddDir(path string) (err error) {
	prefix := filepath.Clean(path)
	return filepath.Walk(prefix, func(path string, fi os.FileInfo, e error) (err error) {
		if e != nil {
			log.Println("errIn", e)
		}

		var filename string
		var data []byte

		if filename, err = filepath.Rel(prefix, path); err != nil {
			log.Println("filename", err)
			return
		}

		if filename == "." {
			filename = "/"
		} else {
			filename = "/" + filename
		}

		// Only read regular types
		if fi.Mode()&os.ModeType == 0 {
			data, err = ioutil.ReadFile(path)
		}

		self.Add(filename, fi, data, err)

		return
	})
}

func (self MemoryFileSystem) Add(name string, fi os.FileInfo, data []byte, err error) {
	self[name] = &MemoryFile{&self, name, fi.Mode(), fi.ModTime(), data, err}
	log.Println("Add", name)
}

type MemoryFile struct {
	fs      *MemoryFileSystem
	name    string
	mode    os.FileMode
	modTime time.Time
	data    []byte
	err     error
}

func (self *MemoryFile) Name() string {
	return filepath.Base(self.name)
}

func (self *MemoryFile) Size() int64 {
	return int64(len(self.data))
}

func (self *MemoryFile) Mode() os.FileMode {
	return self.mode
}

func (self *MemoryFile) ModTime() time.Time {
	return self.modTime
}

func (self *MemoryFile) IsDir() bool {
	return self.mode&os.ModeDir != 0
}

func (self *MemoryFile) Sys() interface{} {
	return nil
}

type MemRef struct {
	*bytes.Reader
	file *MemoryFile
	pos  int // Readdir pos
}

func (self *MemRef) Close() error {
	// TODO: Make the Read operation fail when closed ?
	return nil
}

func (self *MemRef) Stat() (os.FileInfo, error) {
	return self.file, nil
}

// Readdir reads the contents of the directory associated with file and
// returns a slice of up to n FileInfo values, as would be returned by Lstat,
// in directory order. Subsequent calls on the same file will yield further
// FileInfos.
//
// If n > 0, Readdir returns at most n FileInfo structures. In this case, if
// Readdir returns an empty slice, it will return a non-nil error explaining
// why.  At the end of a directory, the error is io.EOF.
//
// If n <= 0, Readdir returns all the FileInfo from the directory in a single
// slice. In this case, if Readdir succeeds (reads all the way to the end of
// the directory), it returns the slice and a nil error. If it encounters an
// error before the end of the directory, Readdir returns the FileInfo read
// until that point and a non-nil error.
func (self *MemRef) Readdir(n int) (files []os.FileInfo, err error) {
	x := self.pos
	size := n
	if size <= 0 {
		size = 100
		n = -1
	}

	files = make([]os.FileInfo, 0, size) // Empty with room to grow.

	path := self.file.name

	for key, file := range *self.file.fs {
		if x != 0 {
			x = -1
			continue
		}

		self.pos += 1
		x, _ := filepath.Rel(path, key)
		if !(strings.HasPrefix(x, ".") || strings.Contains(x, "/")) {
			files = append(files, os.FileInfo(file))
			n -= 1

			if n == 0 {
				break
			}
		}
	}

	if n >= 0 && len(files) == 0 {
		err = io.EOF
	}

	return
}
