package main

import "github.com/kabbythegreat/fileop"

func main() {
	fileop.WriteFileWithValue("file.json", "This is a new file")
}
