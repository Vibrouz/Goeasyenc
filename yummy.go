//go:build ignore

package main

import (
    "io"
    "os"
    "bytes"
    "crypto/rc4"
)


func debomb(key string, filename string) {
    file, err := os.OpenFile(filename, os.O_RDWR, 0664)
    if err != nil { os.Exit(1) }
    defer file.Close()

    var buf bytes.Buffer

    _, err = io.Copy(&buf, file)
    if err != nil { os.Exit(1) }
    file.Seek(0, io.SeekStart)

    trailIndex := buf.Len() - 4
    if !bytes.Equal(buf.Bytes()[trailIndex:], []byte{0xDE, 0xED, 0xBE, 0xEF}) {
        return
    }
    buff := buf.Bytes()[:trailIndex]

    cipher, _ := rc4.NewCipher([]byte(key))
    cipher.XORKeyStream(buff, buff)

    if _, err = file.Write(buff); err != nil { os.Exit(1) }
    if err = file.Truncate(int64(len(buff))); err != nil { os.Exit(1) }
}


func main() {
    debomb("Vibrouz", "file.txt")
}

