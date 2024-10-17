// Embedding Interfaces
package main

import (
	"bytes"
	"fmt"
)

// Reader interface
type Reader interface {
	Read(p []byte) (n int, err error)
}

// Writer interface
type Writer interface {
	Write(p []byte) (n int, err error)
}

// ReaderWriter interface, Reader ve Writer'ı embed ediyor
type ReaderWriter interface {
	Reader
	Writer
}

// MyBuffer struct'ı, hem Reader hem de Writer'ı implemente ediyor
type MyBuffer struct {
	buffer bytes.Buffer
}

// Write metodu: MyBuffer için Write fonksiyonu
func (b *MyBuffer) Write(p []byte) (n int, err error) {
	return b.buffer.Write(p)
}

// Read metodu: MyBuffer için Read fonksiyonu
func (b *MyBuffer) Read(p []byte) (n int, err error) {
	return b.buffer.Read(p)
}

func main() {
	var rw ReaderWriter = &MyBuffer{} // ReaderWriter interface'ini implement eden bir MyBuffer nesnesi oluşturuyoruz

	// Yazma işlemi
	data := []byte("Merhaba Go!")
	n, err := rw.Write(data)
	if err != nil {
		fmt.Println("Yazma hatası:", err)
	}
	fmt.Printf("%d byte yazıldı\n", n)

	// Okuma işlemi
	readBuffer := make([]byte, len(data))
	n, err = rw.Read(readBuffer)
	if err != nil {
		fmt.Println("Okuma hatası:", err)
	}
	fmt.Printf("%d byte okundu: %s\n", n, string(readBuffer))
}
