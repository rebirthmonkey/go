package util

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
)

// WriteFile write data to file system
func WriteFile(filePath string, fileName string, data []byte) error {
	if len(data) == 0 {
		return nil
	}
	if err := EnsurePath(filePath); err != nil {
		return err
	}
	fullName := filepath.Join(filePath, fileName)
	return ioutil.WriteFile(fullName, data, 0666)
}

// AppendFile append data to file system
func AppendFile(filePath string, fileName string, data []byte) error {
	if len(data) == 0 {
		return nil
	}
	var err error
	err = EnsurePath(filePath)
	if err != nil {
		return err
	}
	fullName := filepath.Join(filePath, fileName)
	var file *os.File
	_, err = os.Stat(fullName)
	if os.IsNotExist(err) {
		file, err = os.OpenFile(fullName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	}
	if err != nil {
		return err
	}
	if file == nil {
		if file, err = os.OpenFile(fullName, os.O_APPEND|os.O_WRONLY, 0666); err != nil {
			return err
		}
	}
	defer file.Close()
	_, err = file.Write(data)
	return err
}

// EnsurePath ensure path exist
func EnsurePath(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
	}
	return err
}

func ScanFile(file string, callback func(string)) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		callback(scanner.Text())
	}
	return nil
}

func ListFiles(dirname string) ([]string, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	names, err := f.Readdirnames(-1)
	if err != nil {
		return nil, err
	}
	f.Close()
	sort.Strings(names)
	return names, nil
}
