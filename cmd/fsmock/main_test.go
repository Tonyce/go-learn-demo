package main

import (
	"testing"
)

// func Test_analyze(t *testing.T) {
// 	t.Run("Test something", func(t *testing.T) {
// 		if err := doSomething(strings.NewReader("This is a test string")); (err != nil) != false {
// 			t.Errorf("analyze() error = %v", err)
// 		}
// 	})
// }

func Test_getSize(t *testing.T) {
	oldFs := fs
	// Create and "install" mocked fs:
	mfs := &mockedFS{}
	fs = mfs
	// Make sure fs is restored after this test:
	defer func() {
		fs = oldFs
	}()

	// Test when filesystem.Stat() reports error:
	mfs.reportErr = true
	if _, err := getSize("hello.go"); err == nil {
		t.Error("Expected error, but err is nil!")
	}

	// Test when no error and size is returned:
	mfs.reportErr = false
	mfs.reportSize = 123
	if size, err := getSize("hello.go"); err != nil {
		t.Errorf("Expected no error, got: %v", err)
	} else if size != 123 {
		t.Errorf("Expected size %d, got: %d", 123, size)
	}
}
