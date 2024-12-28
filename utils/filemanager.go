package utils

import (
	"io"
	"os"
)

type FileI interface {
	Close()
	Size() int64
}

type FileMI interface {
	Read() []byte
	Write(data []byte)
}

type File struct {
	File *os.File
}

type FileM struct {
	Directory string
}

func (j File) Close() {
	err := j.File.Close()
	if err != nil {
		panic(err)
	}
}

func FileOpen(file string, flag int, perm os.FileMode) File {
	fi, err := os.OpenFile(file, flag, perm)
	if err != nil {
		panic(err)
	}
	return File{File: fi}
}

func (j *File) Size() int64 {
	fi, err := j.File.Stat()
	if err != nil {
		panic(err)
	}

	return fi.Size()
}

func (j FileM) Read() []byte {
	file := FileOpen(j.Directory, os.O_RDWR|os.O_CREATE, 0755)
	defer file.Close()

	buf := make([]byte, file.Size())

	_, err := file.File.Read(buf)
	if err != nil && err != io.EOF {
		panic(err)
	}

	return buf
}

func (j FileM) Write(data []byte) {
	file := FileOpen(j.Directory, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	defer file.Close()

	_, err := file.File.Write(data)

	if err != nil {
		panic(err)
	}
}
