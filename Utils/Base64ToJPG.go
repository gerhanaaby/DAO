package utils

import (
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"strings"
)

func Base64toJpg(data string, path string, cif string) string {

	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))
	m, formatString, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()
	fmt.Println("base64toJpg", bounds, formatString)

	//Encode from image format to writer
	var jpgFileName string
	if path == "KTP" {
		jpgFileName = "images/KTP/" + cif + ".jpg"
	} else {
		jpgFileName = "images/SelfieKTP/" + cif + ".jpg"
	}

	//jpgFilename := "test.jpg"
	f, err := os.OpenFile(jpgFileName, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}

	err = jpeg.Encode(f, m, &jpeg.Options{Quality: 75})
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}
	fmt.Println("Jpg file", jpgFileName, "created")
	return f.Name()

}
