// +build !windows

package libpcre

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// TestPcre tests the transpiled libpcre using the transpiled pcretest program from the original pcre library.
// The files are copied into a temp directory.
// All .go files are changed to have package main, then pcretest is build.
// Finally RunTest is used to run the tests.
func TestPcre(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "pcre-test")
	if err != nil {
		t.Errorf("Cannot create temp folder: %v", err)
		return
	}
	defer func() { // clean up
		err := os.RemoveAll(tempDir)
		if err != nil {
			t.Errorf("Cannot cleanup: %v", err)
		}
	}()
	srcDir := filepath.Clean("./")
	err = CopyDir(srcDir, tempDir, getShouldCopyFunc(srcDir))
	if err != nil {
		t.Errorf("Cannot copy to temp folder: %v", err)
		return
	}
	err = changeGoPackage2Main(tempDir)
	if err != nil {
		t.Errorf("Cannot change Go package to main: %v", err)
		return
	}
	err = goBuildPcretest(srcDir, tempDir)
	if err != nil {
		t.Errorf("Cannot build pcretest: %v", err)
		return
	}
	err = runPcreTest(t, srcDir)
	if err != nil {
		t.Errorf("Cannot run tests: %v", err)
		return
	}
}

func changeGoPackage2Main(tmpDir string) error {
	files, err := ioutil.ReadDir(tmpDir)
	if err != nil {
		return err
	}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		if ext := path.Ext(f.Name()); ext != ".go" {
			continue
		}
		input, err := ioutil.ReadFile(filepath.Join(tmpDir, f.Name()))
		if err != nil {
			return err
		}

		lines := strings.Split(string(input), "\n")

		for i, line := range lines {
			if strings.Contains(line, "package libpcre") {
				lines[i] = "package main"
				break
			}
		}
		output := strings.Join(lines, "\n")
		err = ioutil.WriteFile(filepath.Join(tmpDir, f.Name()), []byte(output), 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

func goBuildPcretest(srcDir string, tmpDir string) error {
	args := []string{"build", "-tags", "pcretest", "pcretest.go", "libpcre.go", "api.go", "pcre_constants.go", "pcre_added.go", "pcretest_added.go"}
	if runtime.GOOS == "linux" {
		args = append(args, "pcretest_added_linux.go")
	} else if runtime.GOOS == "darwin" {
		args = append(args, "pcretest_added_darwin.go")
	} else if runtime.GOOS == "windows" {
		args = append(args, "pcretest_added_windows.go")
	}
	cmd := exec.Command("go", args...)
	cmd.Dir = tmpDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return CopyFile(filepath.Join(tmpDir, "pcretest"), filepath.Join(srcDir, "pcretest"))
}

func runPcreTest(t *testing.T, srcDir string) error {
	// TODO: fix errors in other tests:
	//    2, 5, 7, 8, 14, 16 printf formatting handling (in c2go)
	//    3, 15 fix locale error
	args := strings.Split("1 4 6 9 10 11 12 13 17 18 19 20 21 22 23 24 25 26", " ")
	cmd := exec.Command("./RunTest", args...)
	cmd.Dir = srcDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err == nil {
		return nil
	}
	if ex, ok := err.(*exec.ExitError); ok {
		t.Errorf("Failed test: %v", ex)
		return nil
	}
	return err
}

// ShouldCopy functions return whether this file/directory should be copied.
type ShouldCopy func(src string, info os.FileInfo) bool

func getShouldCopyFunc(srcDir string) ShouldCopy {
	return func(src string, info os.FileInfo) bool {
		if srcDir == src {
			return true
		}
		ext := path.Ext(src)
		if ext == ".go" {
			return true
		}
		return false
	}
}

// CopyFile copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file. The file mode will be copied from the source and
// the copied data is synced/flushed to stable storage.
func CopyFile(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		if e := out.Close(); e != nil {
			err = e
		}
	}()

	_, err = io.Copy(out, in)
	if err != nil {
		return
	}

	err = out.Sync()
	if err != nil {
		return
	}

	si, err := os.Stat(src)
	if err != nil {
		return
	}
	err = os.Chmod(dst, si.Mode())
	if err != nil {
		return
	}

	return
}

// CopyDir recursively copies a directory tree, attempting to preserve permissions.
// Source directory must exist, destination directory must *not* exist.
// Symlinks are ignored and skipped.
func CopyDir(src string, dst string, shouldCopy ShouldCopy) (err error) {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)

	si, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !si.IsDir() {
		return fmt.Errorf("source is not a directory")
	}
	if !shouldCopy(src, si) {
		return nil
	}

	_, err = os.Stat(dst)
	if err != nil && !os.IsNotExist(err) {
		return
	}
	//if err == nil {
	//	return fmt.Errorf("destination already exists")
	//}

	err = os.MkdirAll(dst, si.Mode())
	if err != nil {
		return
	}

	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			err = CopyDir(srcPath, dstPath, shouldCopy)
			if err != nil {
				return
			}
		} else {
			// Skip symlinks.
			if entry.Mode()&os.ModeSymlink != 0 {
				continue
			}
			if !shouldCopy(srcPath, entry) {
				continue
			}

			err = CopyFile(srcPath, dstPath)
			if err != nil {
				return
			}
		}
	}

	return
}
