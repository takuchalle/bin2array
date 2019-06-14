package main

import (
	"os"
	"log"
	b "github.com/takuyaohashi/bin2array"
)

func main() {
	file, err := os.Open("main.go")
	handleError(err)

	info, _ := file.Stat()
	size := info.Size()
	
	o := &b.Options{OutStream: os.Stdout, InStream: file, Size: size}
	c, err := b.New(o)
	if err != nil {
		panic(err)
	}

	err = c.Run()
	if err != nil {
		panic(err)
	}

	os.Exit(0)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
