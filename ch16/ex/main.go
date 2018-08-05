package main

import (
	"crypto/md5"
	"encoding/hex"
	"crypto/aes"
	"crypto/cipher"
	"io"
	"crypto/rand"
	"fmt"
	"os"
	"io/ioutil"
)

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(block)
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, _ := gcm.Open(nil, nonce, ciphertext, nil)
	return plaintext
}

func encryptFile(filename string, data []byte, passphrase string)  {
	f, _ := os.Create(filename)
	defer f.Close()
	f.Write(encrypt(data, passphrase))
}

func decryptFile(filename string,  passphrase string) []byte {
	data, _ := ioutil.ReadFile(filename)
	return decrypt(data, passphrase)
}


func main() {
	ciphertext := encrypt([]byte("Hello World"), "password")
	fmt.Println(string(ciphertext))
	plaintext := decrypt(ciphertext, "password")
	fmt.Println(string(plaintext))

	encryptFile("example.txt", []byte("Hello World"), "password")
	plaintext = decryptFile("example.txt", "password")
	fmt.Println(string(plaintext))
}
