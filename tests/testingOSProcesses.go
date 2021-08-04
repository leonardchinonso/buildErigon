package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	dir, err := os.MkdirTemp("C:\\Users\\USER\\Documents\\Learning\\testingGo", "example")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			log.Fatal(fmt.Errorf("cannot remove this folder %s", err))
		}
	}(dir) // clean up

	file := filepath.Join(dir, "tempFile")
	if err := os.WriteFile(file, []byte("content"), 0666); err != nil {
		log.Fatal(err)
	}
}
