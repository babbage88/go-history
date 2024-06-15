package main

import (
	"fmt"
	"os"
	"time"

	fileops "github.com/babbage88/go-history/utils"
)

func main() {

	start := time.Now()
	scanhistory, err := fileops.SearchByLine("export")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	elapsed := time.Since(start)
	defer fmt.Printf("page took %s\n", elapsed)

	fmt.Println(len(scanhistory))
	//fmt.Println(len(history))

}
