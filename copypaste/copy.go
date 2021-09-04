package copypaste

import (
	"errors"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func createFolder(folderName string) {
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		os.MkdirAll(folderName, os.ModePerm)
	}
}

func copy(fileName string, copyFile string) (error, error) {
	f1, err1 := os.Open(fileName)
	f2, err2 := os.Create(copyFile)
	defer func() { f1.Close(); f2.Close() }()
	io.Copy(f2, f1)
	return err1, err2
}

func ListAllPaths(root string) ([]string, []string) {
	var pathFiles []string
	var pathDir []string

	var err error = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		var isdir bool = info.IsDir()

		if isdir {
			pathDir = append(pathDir, path)
		} else {
			pathFiles = append(pathFiles, path)
		}
		return nil
	})
	check(err)
	return pathDir, pathFiles
}

func checkError(err error) bool {
	if err != nil {
		return true
	}
	return false
}

func CopyFolder(pathDir []string, pathFiles []string, old string, nouv string) error {
	for _, dir := range pathDir {
		newName := strings.Replace(dir, old, nouv, 1)
		createFolder(newName)
	}

	for _, filename := range pathFiles {
		newname := strings.Replace(filename, old, nouv, 1)
		err1, err2 := copy(filename, newname)
		if checkError(err1) || checkError(err2) {
			return errors.New("Problème")

		}
	}
	return nil
}

func CopyPaste(oldFolder string, newFolder string) error {
	root := strings.Split(oldFolder, "/")
	newFolder = path.Join(newFolder, root[len(root)-1])
	pathsDir, pathsFiles := ListAllPaths(oldFolder)
	Err := CopyFolder(pathsDir, pathsFiles, oldFolder, newFolder)
	return Err
}
