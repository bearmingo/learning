package main

import "bytes"

func main() {
	// code_error()
	code()
}

func code_error() {
	path := []byte("AAAA/BBBBBBBB")
	sep := bytes.IndexByte(path, '/')
	println("sep: ", sep) // sep:  4

	dir1 := path[:sep]
	dir2 := path[sep+1:]
	println("dir1: ", string(dir1)) // dir1:  AAAA
	println("dir2: ", string(dir2)) // dir2:  BBBBBBBB

	dir1 = append(dir1, "suffix"...)
	println("current dir1: ", string(dir1)) // current dir1:  AAAAsuffix

	path = bytes.Join([][]byte{dir1, dir2}, []byte{'/'})
	println("dir1: ", string(dir1)) // dir1:  AAAAsuffix
	println("dir2: ", string(dir2)) // dir2:  uffixBBB

	println("new path: ", string(path)) // new path:  AAAAsuffix/uffixBBB
}

// 使用 full slice expression
func code() {
	path := []byte("AAAA/BBBBBBBB")
	sep := bytes.IndexByte(path, '/')
	println("sep: ", sep) // sep:  4

	dir1 := path[:sep:sep] // 此时 cap(dir1) 指定为4， 而不是先前的 16
	dir2 := path[sep+1:]

	path = bytes.Join([][]byte{dir1, dir2}, []byte{'/'})
	println("dir1: ", string(dir1))     // dir1:  AAAAsuffix
	println("dir2: ", string(dir2))     // dir2:  BBBBBBBB
	println("new path: ", string(path)) // new path:  AAAA/BBBBBBBB
}
