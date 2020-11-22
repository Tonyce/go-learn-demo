package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

var fs fileSystem = osFS{}

type fileSystem interface {
	Open(name string) (file, error)
	Stat(name string) (os.FileInfo, error)
}

type file interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	Stat() (os.FileInfo, error)
}

// osFS implements fileSystem using the local disk.
type osFS struct{}

func (osFS) Open(name string) (file, error)        { return os.Open(name) }
func (osFS) Stat(name string) (os.FileInfo, error) { return os.Stat(name) }

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getSize(name string) (int64, error) {
	stat, err := fs.Stat(name)
	if err != nil {
		return 0, err
	}
	return stat.Size(), nil
}

func main() {
	analyze("test.txt")

	dat, err := ioutil.ReadFile("cmd/fsmock/refer.md")
	check(err)
	fmt.Print(string(dat))

	size, err := getSize("cmd/fsmock/refer.md")
	check(err)
	fmt.Print(size)
}

func analyze(file string) error {
	handle, err := os.Open(file)

	if err != nil {
		return err
	}
	defer handle.Close()
	return doSomething(handle)
}

func doSomething(handle io.Reader) error {
	scanner := bufio.NewScanner(handle)
	for scanner.Scan() {
		// Do something with line
		_ = scanner.Text()
	}
	return nil
}
