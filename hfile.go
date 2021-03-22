package hfile

import (
	"crypto/sha256"
	"fmt"
	"io"
	"math/rand"
	"os"
)

const (
	SecSize = 4 << 10
)

func Hash(fileName string) (result []byte, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer f.Close()
	fs, err := f.Stat()
	if err != nil {
		return
	}
	sectionReader := io.NewSectionReader(f, 0, fs.Size())
	hash256 := sha256.New()
	hash256.Write([]byte(fmt.Sprint(sectionReader.Size())))
	sz := sectionReader.Size()
	if sz < SecSize {
		buff := make([]byte, sz)
		sectionReader.Read(buff)
		hash256.Write(buff)
	} else {
		buff := make([]byte, SecSize)
		sectionReader.Read(buff)
		hash256.Write(buff)
		carry := sz / SecSize
		rand.Seed(carry)
		rd := rand.Int63n(sz - SecSize)
		sectionReader.Seek(rd, 1)
		sectionReader.Read(buff)
		hash256.Write(buff)
	}
	result = hash256.Sum(nil)
	return
}

func HashString(fileName string) (result string, err error) {
	sum, err := Hash(fileName)
	if err != nil {
		return
	}
	result = fmt.Sprintf("%x", sum)
	return
}
