package main

import (
	"USBDetection/copypaste"
	"os"
	"path"
	"strings"
	"time"
)

func contains(list []string, x string) bool {
	for _, element := range list {
		if element == x {
			return true
		}
	}
	return false
}

func substring(fullkeyname string, sub string) bool {
	return strings.Contains(fullkeyname, sub)
}

func main() {
	var usbPath string = copypaste.LocationUSB("/media")
	var keyDetected []string
	cwd, err := os.Getwd()
	cwd = path.Join(cwd, "clés")

	if err != nil {
		panic(err)
	}

	for {
		listKey := copypaste.DetectionUSB(usbPath)
		if len(listKey) != 0 {
			for _, key := range listKey {
				if !contains(keyDetected, key) {
					keyDetected = append(keyDetected, key)
					go func() { copypaste.CopyPaste(key, cwd) }()
				}
			}
		}
		time.Sleep(2 * time.Second)
	}
}
