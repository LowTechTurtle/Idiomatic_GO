package cleanup

import (
	"errors"
	"os"
	"strings"
	"testing"
)

func createFile(t *testing.T) (string, error) {
	f, err := os.Create("tempFile")
	if err != nil {
		return "", err
	}
	defer func() {
		err = errors.Join(err, f.Close())
	}()

	_, errw := f.Write([]byte("banana"))
	if errw != nil {
		err = errors.Join(err, errw)
	}

	t.Cleanup(func() {
		os.Remove(f.Name())
	})
	return f.Name(), nil
}

func TestFileProcessing(t *testing.T) {
	fName, err := createFile(t)
	if err != nil {
		t.Fatal(err)
	}
	// do testing, don't worry about cleanup
	if !strings.Contains(fName, "tempFile") {
		t.Error("unexpected name")
	}
}

func createFileTempDir(tempDir string) (string, error) {
	f, err := os.CreateTemp(tempDir, "tempFile")
	if err != nil {
		return "", err
	}
	defer func() {
		err = errors.Join(err, f.Close())
	}()

	_, errw := f.Write([]byte("banana"))
	if errw != nil {
		err = errors.Join(err, errw)
	}
	return f.Name(), nil
}

func TestFileProcessingTempDir(t *testing.T) {
	tempDir := t.TempDir()
	fname, err := createFileTempDir(tempDir)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(fname, "tempFile") {
		t.Error("unexpected name")
	}
}
