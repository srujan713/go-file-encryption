package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/akhilsharma90/go-file-encryption/filecrypt"
	"golang.org/x/term"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}
	function := os.Args[1]

	switch function {
	case "help":
		printHelp()
	case "encrypt":
		encryptHandle()
	case "decrypt":
		decryptHandle()
	default:
		fmt.Println("Run enrypt to encrypt a file, and decrypt to decrypt a file.")
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println("File Encryption:")
	fmt.Println("Simple File Encrypter for day-to-day tasks:")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("\t go run . encrypt /path/to/your/file")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("")
	fmt.Println("\t encrypt \t Encrypt a file given a password")
	fmt.Println("\t decrypt \t Decrypt a file using a password")
	fmt.Println("\t help \t\t Display help text")
	fmt.Println("")
}

func encryptHandle() {
	if len(os.Args) < 3 {
		fmt.Println("Missing the path to the file. For more info, run go run . help")
		os.Exit(0)
	}
	file := os.Args[2]
	if !validateFile(file) {
		panic("File not found")
	}
	password := getPassword()
	fmt.Println("\nEncrypting...")
	filecrypt.Encrypt(file, password)
	fmt.Println("\n file successfully protected")
}

func decryptHandle() {
	if len(os.Args) < 3 {
		fmt.Println("Missing the path to the file. For more info, run  go run . help")
		os.Exit(0)
	}
	file := os.Args[2]
	if !validateFile(file) {
		panic("File not found")
	}
	fmt.Print("Enter password:")
	password, _ := term.ReadPassword(0)
	fmt.Println("\nDecrypting...")
	filecrypt.Decrypt(file, password)
	fmt.Println("\n file successfully protected")
}

func getPassword() []byte {
	fmt.Print("Enter password")
	password, _ := term.ReadPassword(0)
	fmt.Print("\nConfirm Password: ")
	password2, _ := term.ReadPassword(0)
	if !validatePassword(password, password2) {
		fmt.Print("\nPassword do not match. Please Try Again...\n")
		return getPassword()
	}
	return password
}

func validatePassword(password1 []byte, password2 []byte) bool {
	if !bytes.Equal(password1, password2) {
		return false
	}
	return true
}

func validateFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}
