package templateengine

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)
// CopyFile Copies Files around
func CopyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			_ = os.Chmod(dest, sourceinfo.Mode())
		}

	}

	return
}

//IsDirectory checks if
func IsDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	return fileInfo.IsDir(), err
}
//ListDir Lists files in dir
func ListDir(dir string) []os.FileInfo {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Println(" -", file.Name())
		}
	}

	return files
}
//ListDirWithExceptions Lists files in dir, skipping elements in array list
func ListDirWithExceptions(dir string, exceptions []string) []os.FileInfo {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		errorResp := []os.FileInfo{}
		return errorResp
	}

	for _, file := range files {
		if file.IsDir() {
			if !StringInSlice(file.Name(), exceptions) {
				fmt.Println(file.Name())
			}
		}
	}

	return files
}
//CopyDir Copies directories
func CopyDir(source string, dest string) (err error) {

	// get properties of source dir
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// create dest dir

	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		sourcefilepointer := source + "/" + obj.Name()

		destinationfilepointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			// create sub-directories - recursively
			err = CopyDir(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			// perform copy
			err = CopyFile(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	return
}

//Exists checks if
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

//CreateNewDirectoryIfNil checks and creates fi needed
func CreateNewDirectoryIfNil(path string) (bool, error) {
	exists, err := Exists(path)
	if err != nil {
		return false, err
	}
	if !exists {
		os.MkdirAll(path, 00750)
		return true, nil
	}
	return false, nil
}
