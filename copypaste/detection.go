package copypaste

import (
	"io/ioutil"
	"path"
)

// LocationUSB : retourne le chemin avec le nom de la session d'utilisateur pour accéder aux clés USB
func LocationUSB(pathMedia string) string {
	listUser, err := ioutil.ReadDir(pathMedia)
	check(err)
	return path.Join(pathMedia, listUser[0].Name())
}

// DetectionUSB : renvoie les dossiers à copier coller
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
