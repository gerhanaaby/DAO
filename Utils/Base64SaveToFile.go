package utils

import (
	"encoding/base64"
	"os"
)

func Base64toFile(data string, path string, cif string) string {
	dec, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		panic(err)
	}

	var base64FileName string
	if path == "KTP" {
		base64FileName = "images/KTP/" + cif
	} else {
		base64FileName = "images/SelfieKTP/" + cif
	}

	f, err := os.Create(base64FileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		panic(err)
	}
	if err := f.Sync(); err != nil {
		panic(err)
	}
	return base64FileName
}
