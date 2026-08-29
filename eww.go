//go:build ignore

package main

import (
	"io"
	"os"
	"bytes"
	"crypto/rc4"
)


func bomb(key string, filename string) {
    file, err := os.OpenFile(filename, os.O_RDWR, 0664)
    if err != nil { os.Exit(1) }
    defer file.Close()

    var buf bytes.Buffer

    _, err = io.Copy(&buf, file)
    if err != nil { os.Exit(1)}
    file.Seek(0, io.SeekStart)

    if bytes.Equal(buf.Bytes()[buf.Len()-4:], []byte{0xDE, 0xED, 0xBE, 0xEF}) {
        return
    }

    cipher, _ := rc4.NewCipher([]byte(key))
    cipher.XORKeyStream(buf.Bytes(), buf.Bytes())

    if _, err := file.Write(buf.Bytes()); err != nil { os.Exit(1) }
    if _, err := file.Write([]byte{0xDE, 0xED, 0xBE, 0xEF}); err != nil { os.Exit(1) }
}


func main() {
    bomb("Vibrouz", "file.txt")
}


