package copypaste

import (
	"io/ioutil"
	"path"
)

func LocationUSB(pathMedia string) string {
	listUser, err := ioutil.ReadDir(pathMedia)
	check(err)
	return path.Join(pathMedia, listUser[0].Name())
}

func DetectionUSB(keyPath string) []string {
	var ListKeyPaths []string
	KeyLists, err := ioutil.ReadDir(keyPath)
	check(err)
	if len(KeyLists) != 0 {
		for _, key := range KeyLists {
			ListKeyPaths = append(ListKeyPaths, path.Join(keyPath, key.Name()))
		}
	}
	return ListKeyPaths

}
