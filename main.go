package main

import (
	"crypto/md5"
	"encoding/hex"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	win := myApp.NewWindow("Text Encryptor")

	// Set fixed size for the window (width, height)1
	win.Resize(fyne.NewSize(400, 300))
	win.SetFixedSize(true)

	inputEntry := widget.NewEntry()
	MD5Label := widget.NewLabel("MD5:")
	EncryptedLabel := widget.NewLabel("Encrypted:")
	var MD5Output string
	var EncryptedOutput string
	copyEncryptedButton := widget.NewButton("Copy", func() {
		clipboard := win.Clipboard()
		if clipboard == nil {
			return
		}
		clipboard.SetContent(EncryptedOutput)
	})
	
	copyMD5Button := widget.NewButton("Copy", func() {
		clipboard := win.Clipboard()
		if clipboard == nil {
			return
		}
		clipboard.SetContent(MD5Output)
	})

	encryptButton := widget.NewButton("Encrypt", func() {
		input := inputEntry.Text
		encryptedStr := encryptText(input)
		hash := encryptAndHash(encryptedStr)
		EncryptedOutput = encryptedStr
		MD5Output = hash
		MD5Label.SetText("MD5: " + MD5Output)
		EncryptedLabel.SetText("Encrypted: " + EncryptedOutput)
	})

	// Horizontal layout with spacing between elements
	encryptedRow := container.NewHBox(
		EncryptedLabel,
		layout.NewSpacer(),
		copyEncryptedButton,
	)

	md5Row := container.NewHBox(
		MD5Label,
		layout.NewSpacer(),
		copyMD5Button,
	)

	content := container.NewVBox(
		widget.NewLabel("Enter text to encrypt:"),
		inputEntry,
		encryptButton,
		encryptedRow,
		md5Row,
	)
	win.SetContent(content)
	win.ShowAndRun()
}

func encryptText(input string) string {
	input = strings.TrimSpace(input)
	encrypted := make([]int, len(input))
	for i, c := range input {
		encrypted[i] = int(c) * 2
		for encrypted[i] > 126 {
			encrypted[i] -= 94
		}
		if encrypted[i] == 32 {
			encrypted[i] += 1
		}
	}
	var encryptedStr string
	for _, v := range encrypted {
		encryptedStr += string(rune(v))
	}
	return encryptedStr
}

func encryptAndHash(text string) string {
	hash := md5Hash(text)
	for i := 0; i < 3; i++ {
		hash = md5Hash(hash)
	}
	return hash
}

func md5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
