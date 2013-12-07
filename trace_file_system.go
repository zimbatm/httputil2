package httputil2

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

// Used to debug http.FileSystem interfaces.
func NewTraceFileSystem(f http.FileSystem) http.FileSystem {
	return &traceFS{f}
}

type traceFS struct {
	http.FileSystem
}

func (t *traceFS) Open(name string) (f http.File, e error) {
	f, e = t.FileSystem.Open(name)
	fmt.Println("Open", name, f, e)
	f = &traceFile{f, name}
	return
}

type traceFile struct {
	http.File
	name string
}

func (t *traceFile) Close() (e error) {
	e = t.File.Close()
	fmt.Println("File Close", t.name, e)
	return
}

func (t *traceFile) Stat() (fi os.FileInfo, e error) {
	fi, e = t.File.Stat()
	fmt.Println("File Stat", t.name, fi, e)
	fi = &traceFileInfo{fi, t.name}
	return
}

func (t *traceFile) Readdir(count int) (fis []os.FileInfo, e error) {
	fis, e = t.File.Readdir(count)
	fmt.Println("File Readdir", t.name, count, fis, e)
	return
}

func (t *traceFile) Read(data []byte) (n int, e error) {
	n, e = t.File.Read(data)
	fmt.Println("File Read", t.name, data, n, e)
	return
}

func (t *traceFile) Seek(offset int64, whence int) (n int64, e error) {
	n, e = t.File.Seek(offset, whence)
	fmt.Println("File Seek", t.name, offset, whence, n, e)
	return
}

type traceFileInfo struct {
	os.FileInfo
	name string
}

func (t *traceFileInfo) Name() (x string) {
	x = t.FileInfo.Name()
	fmt.Println("FileInfo Name", t.name, x)
	return
}

func (t *traceFileInfo) Size() (x int64) {
	x = t.FileInfo.Size()
	fmt.Println("FileInfo Size", t.name, x)
	return
}

func (t *traceFileInfo) Mode() (x os.FileMode) {
	x = t.FileInfo.Mode()
	fmt.Println("FileInfo Mode", t.name, x)
	return
}

func (t *traceFileInfo) ModTime() (x time.Time) {
	x = t.FileInfo.ModTime()
	fmt.Println("FileInfo ModTime", t.name, x)
	return
}

func (t *traceFileInfo) IsDir() (x bool) {
	x = t.FileInfo.IsDir()
	fmt.Println("FileInfo IsDir", t.name, x)
	return
}

func (t *traceFileInfo) Sys() (x interface{}) {
	x = t.FileInfo.Sys()
	fmt.Println("FileInfo Sys", t.name, x)
	return
}
