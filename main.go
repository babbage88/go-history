package main

import (
	"fmt"
	"os"
	"time"

	fileops "github.com/babbage88/go-history/utils/fileops"
)

func main() {

	start := time.Now()
	scanhistory, err := fileops.SearchCmdHistory("export")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	elapsed := time.Since(start)
	defer fmt.Printf("page took %s\n", elapsed)

	fmt.Println(len(scanhistory))
	//fmt.Println(len(history))

}
