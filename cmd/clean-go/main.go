package main

import (
	"os"
)

func main() {
	if err := buildCmd().Run(os.Args); err != nil {
		panic(err)
	}
}
