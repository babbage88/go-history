package main

import (
	"fmt"
	"os"
	"time"

	database "github.com/babbage88/go-history/database"
	fileops "github.com/babbage88/go-history/utils/fileops"
)

func devDebug(commands []fileops.CommandHistoryEntry, count int, duration time.Duration, srch *string) {
	//_, countst, durationst := testString(searchPtr, regExPtr)

	for index, line := range commands {
		var previous fileops.CommandHistoryEntry
		if index > 1 {
			previous = commands[index-1]
		}

		fmt.Printf("Previous Entry: Line: %d, BaseCommand: %s\n", previous.LineNumber, previous.BaseCommand)
		fmt.Printf("Index: %d\n", index)
		fmt.Printf("%d ", line.LineNumber)
		fmt.Printf("%s ", line.DateExecuted.Format("2006-1-2 15:04:05 "))
		fmt.Printf("%s ", line.BaseCommand)
		fmt.Printf("%s \n", line.SubCommand)

	}
	fmt.Printf("search for %s struct returned took %s returning %d records\n", *srch, duration, count)
}

func testStruct(srch *string, reg *bool) ([]fileops.CommandHistoryEntry, int, time.Duration) {
	start := time.Now()
	defer fmt.Printf("Testing test: %v\n", *srch)
	scanhistory, err := fileops.GetCmdHistory(*srch, *reg)
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

	//searchPtr := flag.String("srch", ".", "Regex ecpression to search ~/.bash_history")
	//regExPtr := flag.Bool("usereg", false, "Use a regex expressiong for search pattern, can hurt performance.")

	//flag.Parse()

	//_, countst, durationst := testString(searchPtr, regExPtr)
	//retval, count, duration := testStruct(searchPtr, regExPtr)
	//fmt.Println(retval)

	//defer fmt.Printf("se %s struct returned took %s returning %d records\n", *searchPtr, duration, count)
	dbsql := database.NewDatabaseConnection()
	db, _ := database.InitializeDbConnection(dbsql)
	//database.InsertCommandHistoryEntries(db, retval)
	start := time.Now()
	testget, _ := database.GetAllCmdHistory(db)
	elapsed := time.Since(start)
	defer fmt.Printf("The SQL lite query from GetAllCmHistory(db) took %s to complete.", elapsed)

	//fmt.Printf("%v\n", testget)
	fmt.Printf("%d\n", len(testget))
	//defer fmt.Printf("search for %s with string returned took %s returning %d records\n", *searchPtr, durationst, countst)
}
