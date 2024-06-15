package main

import (
	"fmt"
	"os"
	"time"

	fileops "github.com/babbage88/go-history/utils"
)

func main() {

	start := time.Now()
	history, err := fileops.GetFileConcurrent(100)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	elapsed := time.Since(start)

	fmt.Println(history)
	//fmt.Println(len(history))

	fmt.Printf("page took %s\n", elapsed)
}
