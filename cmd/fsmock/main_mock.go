package main

import (
	"os"
)

type mockedFS struct {
	// Embed so we only need to "override" what is used by testable functions
	osFS

	reportErr  bool  // Tells if this mocked FS should return error in our tests
	reportSize int64 // Tells what size should Stat() report in our test
}

type mockedFileInfo struct {
	// Embed this so we only need to add methods used by testable functions
	os.FileInfo
	size int64
}

func (m mockedFileInfo) Size() int64 { return m.size }

func (m mockedFS) Stat(name string) (os.FileInfo, error) {
	if m.reportErr {
		return nil, os.ErrNotExist
	}
	return mockedFileInfo{size: m.reportSize}, nil
}
