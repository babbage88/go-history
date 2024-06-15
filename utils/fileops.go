package fileops

import (
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"sync"
)

type chunk struct {
	bufsize int
	offset  int64
}

func getCurrentUserBashHistoryPath() (string, error) {
	// Get the current user's home directory
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}
	homeDir := currentUser.HomeDir

	historyFilePath := filepath.Join(homeDir, ".bash_history")

	return historyFilePath, nil
}

func GetFileConcurrent(buffsize int) ([]string, error) {
	output := make([]string, 0)

	historyFilePath, err := getCurrentUserBashHistoryPath()
	if err != nil {
		fmt.Errorf("Error finding bash_history path", err)
		return output, err
	}
	file, err := os.Open(historyFilePath)
	if err != nil {
		fmt.Errorf("Error opening file", err)
		return output, err
	}

	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Errorf("Error stating file. Check permissions", err)
		return output, err
	}

	filesize := int(fileinfo.Size())
	// Number of go routines we need to spawn.
	concurrency := filesize / buffsize
	chunksizes := make([]chunk, concurrency)

	// All buffer sizes are the same in the normal case. Offsets depend on the
	// index. Second go routine should start at 100, for example, given our
	// buffer size of 100.
	for i := 0; i < concurrency; i++ {
		chunksizes[i].bufsize = buffsize
		chunksizes[i].offset = int64(buffsize * i)
	}

	// check for any left over bytes. Add one more go routine if required.
	if remainder := filesize % buffsize; remainder != 0 {
		c := chunk{bufsize: remainder, offset: int64(concurrency * buffsize)}
		concurrency++
		chunksizes = append(chunksizes, c)
	}

	var wg sync.WaitGroup
	wg.Add(concurrency)

	for i := 0; i < concurrency; i++ {
		go func(chunksizes []chunk, i int) {
			defer wg.Done()

			chunk := chunksizes[i]
			buffer := make([]byte, chunk.bufsize)
			bytesread, err := file.ReadAt(buffer, chunk.offset)

			if err != nil && err != io.EOF {
				fmt.Println(err)
				return
			}

			output = append(output, string(buffer[:bytesread]))
		}(chunksizes, i)
	}

	wg.Wait()

	return output, nil
}

func GetBashHistoryChunk(buffSize int32) ([]string, error) {
	historyFilePath, err := getCurrentUserBashHistoryPath()
	output := make([]string, 0)
	if err != nil {
		fmt.Errorf("Error finding bash_history path", err)
		return output, err
	}

	file, err := os.Open(historyFilePath)
	if err != nil {
		fmt.Errorf("Error Opening file", err)
		return output, err
	}
	defer file.Close()

	buffer := make([]byte, buffSize)

	for {
		bytesread, err := file.Read(buffer)

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}

			break
		}

		output = append(output, string(buffer[:bytesread]))
	}

	return output, nil
}

func GetBashHistory() (string, error) {

	historyFilePath, err := getCurrentUserBashHistoryPath()
	if err != nil {
		return "", err
	}

	content, err := os.ReadFile(historyFilePath)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
