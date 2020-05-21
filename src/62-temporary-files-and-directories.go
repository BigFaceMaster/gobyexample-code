package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	f, err := ioutil.TempFile("", "sample")
	checkErr(err)

	fmt.Println("Temp file name:", f.Name())

	defer os.Remove(f.Name())

	_, err = f.Write([]byte{1, 2, 3, 4})
	checkErr(err)

	dname, err := ioutil.TempDir("", "sampledir")
	checkErr(err)
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file1")
	err = ioutil.WriteFile(fname, []byte{1, 2}, 0666)
	checkErr(err)
}