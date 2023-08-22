package main

import (
	"ChromeExtensionGenerator/files"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var FOLDERS []string = []string{
	"css",
	"js",
	"html",
	"logo",
}

var EMPTYFILES []string = []string{
	"css/popup.css",
	"css/style.css",
	"js/background.js",
	"js/popup.js",
	"js/script.js",
}

func main() {
	fmt.Print("The files will be downloaded to the folder where you are located. Do you agree? (Y/n) ")
	reader_confirm := bufio.NewReader(os.Stdin)
	confirm, err := reader_confirm.ReadString('\n')
	if err != nil || strings.ToLower(string(confirm[0])) != "y" {
		fmt.Println("Canceled!")
		return
	}

	fmt.Println("Creating folders...")
	for _, foldername := range FOLDERS {
		err := os.Mkdir(foldername, 0777)

		if err != nil {
			fmt.Println("Folder creation error " + foldername + ": " + err.Error())
			return
		}
	}

	fmt.Println("Creating empty files...")
	for _, emptyfile := range EMPTYFILES {
		err := os.WriteFile(emptyfile, []byte(""), 0777)

		if err != nil {
			fmt.Println("File creation error " + emptyfile + ": " + err.Error())
			return
		}
	}

	fmt.Println("Creating files from memory...")
	err = os.WriteFile("html/popup.html", files.Popup_html, 0777)
	if err != nil {
		fmt.Println("File creation error html/popup.html: " + err.Error())
		return
	}

	err = os.WriteFile("logo/logo-16.png", files.Logo_16_png, 0777)
	if err != nil {
		fmt.Println("File creation error logo/logo-16.png: " + err.Error())
		return
	}

	err = os.WriteFile("logo/logo-48.png", files.Logo_48_png, 0777)
	if err != nil {
		fmt.Println("File creation error logo/logo-48.png: " + err.Error())
		return
	}

	err = os.WriteFile("logo/logo-128.png", files.Logo_128_png, 0777)
	if err != nil {
		fmt.Println("File creation error logo/logo-128.png: " + err.Error())
		return
	}

	err = os.WriteFile("manifest.json", files.Manifest_json, 0777)
	if err != nil {
		fmt.Println("File creation error manifest.json: " + err.Error())
		return
	}

	fmt.Println("Success!")
}
