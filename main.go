package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	fileops "github.com/babbage88/go-history/utils/fileops"
)

func testStruct(srch *string, reg *bool) ([]fileops.CommandHistoryEntry, int, time.Duration) {
	start := time.Now()
	scanhistory, err := fileops.SearchCmdHistory(*srch, *reg)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	cnt := len(scanhistory)
	elapsed := time.Since(start)
	return scanhistory, cnt, elapsed
}

func testString(srch *string, reg *bool) ([]string, int, time.Duration) {
	start := time.Now()
	byline, err := fileops.SearchByLine(*srch, *reg)
	if err != nil {
		fmt.Errorf("Error parsing history", err)
	}
	cnt := len(byline)
	elapsed := time.Since(start)

	return byline, cnt, elapsed
}

func main() {
	searchPtr := flag.String("srch", ".", "Regex ecpression to search ~/.bash_history")
	regExPtr := flag.Bool("usereg", false, "Use a regex expressiong for search pattern, can hurt performance.")

	flag.Parse()

	_, countst, durationst := testString(searchPtr, regExPtr)
	retval, count, duration := testStruct(searchPtr, regExPtr)

	for index, line := range retval {
		fmt.Printf("Index: %d\n", index)
		fmt.Printf("%d ", line.LineNumber)
		fmt.Printf("%s ", line.DateExecuted.Format("2006-1-2 15:04:05 "))
		fmt.Printf("%s ", line.BaseCommand)
		fmt.Printf("%s \n", line.SubCommand)

	}

	defer fmt.Printf("search for %s struct returned took %s returning %d records\n", *searchPtr, duration, count)

	defer fmt.Printf("search for %s with string returned took %s returning %d records\n", *searchPtr, durationst, countst)

	//fmt.Println(len(history))

}
