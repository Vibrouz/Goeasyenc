//go:build ignore

package main

import (
	"bytes"
	"crypto/rc4"
	"io"
	"os"
)


func mini_nuclear(key string, filename string) {
    file, err := os.OpenFile(filename, os.O_RDWR, 0664)
    if err != nil { os.Exit(1) }
    defer file.Close()

    var buf bytes.Buffer

    _, err = io.Copy(&buf, file)
    if err != nil { os.Exit(1)}
    file.Seek(0, io.SeekStart)

    cipher, _ := rc4.NewCipher([]byte(key))
    cipher.XORKeyStream(buf.Bytes(), buf.Bytes())

    if _, err := file.Write(buf.Bytes()); err != nil { os.Exit(1) }
}


func main() {
    mini_nuclear("Vibrouz", "file.txt")
}

