package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mholt/archiver"
)

func fillString(retunString string, toLength int) string {
	for {
		lengtString := len(retunString)
		if lengtString < toLength {
			retunString = retunString + ":"
			continue
		}
		break
	}
	return retunString
}

func Unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}

func UnzipNClean(fileToUnzip string) {

	pathString, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	files, _ := ioutil.ReadDir(pathString)
	fmt.Println(len(files))

	fmt.Println("File to Unzip's name is.... : " + fileToUnzip)

	//Unzip the folder
	errr := archiver.Zip.Open(fileToUnzip, "")

	if errr != nil {
		panic(errr)
	}

	//Delete the folder
	os.Remove(fileToUnzip)

	fmt.Println("Deleted the old file")

}

func prepareRendering(file string) {

	fmt.Println("Unzipping the file.... " + file)
	// UnzipNClean(file)
	//Do I need to clean the file name of .zip?
}

func startRendering(file string) {

}

func sendRendering(file string) {

}
