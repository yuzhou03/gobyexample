// Writing files in Go follows similar patterns to the
// ones we saw earlier for reading.

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"os/exec"
	"strings"
	"errors"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// check if path exits
func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// 获得当前程序路径
func GetCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}

	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}

	if i < 0 {
		return "", errors.New(`error: Can't' find "/" or "\".`)
	}
	return string(path[0:i+1]), nil
}

func main() {

	// To start, here's how to dump a string (or just
	// bytes) into a file.
	d1 := []byte("hello\ngo\n")

	// touch file if not exist
	//outfile1 := "/Users/ad/Code/gopath/src/git.tvblack.com/tvblack/gobyexample/out/" + "data1.txt"
	cwd, _ := GetCurrentPath()
	fmt.Println(cwd)
	outfile1 := cwd + "/data1.txt"
	if !Exists(outfile1) {

	}

	err := ioutil.WriteFile(outfile1, d1, 0644)
	check(err)

	os.Exit(0)

	// For more granular writes, open a file for writing.
	f, err := os.Create("./out/data2.txt")
	check(err)

	// It's idiomatic to defer a `Close` immediately
	// after opening a file.
	defer f.Close()

	// You can `Write` byte slices as you'd expect.
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// A `WriteString` is also available.
	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	// Issue a `Sync` to flush writes to stable storage.
	f.Sync()

	// `bufio` provides buffered writers in addition
	// to the buffered readers we saw earlier.
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	// Use `Flush` to ensure all buffered operations have
	// been applied to the underlying writer.
	w.Flush()

}
