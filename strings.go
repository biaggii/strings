package main

import (
	"os"
)

func main() {
	strings := os.Args[1:]
	for _, s := range strings {
		file, err := os.Open(s)
		if err != nil {
			os.Exit(1)
		}
		defer file.Close()
		var buf [512]byte
		for {
			n, err := file.Read(buf[:])
			if err != nil {
				break
			}
			os.Stdout.Write(buf[:n])
		}
	}
}
