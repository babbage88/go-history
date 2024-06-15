package main

import (
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
)

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

func getBashHistoryChunk(buffSize int32) ([]string, error) {
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

		fmt.Println("bytes read: ", bytesread)
		fmt.Println("bytestream to string: ", string(buffer[:bytesread]))
		output = append(output, string(buffer[:bytesread]))
	}

	return output, nil
}

func getBashHistory() (string, error) {

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

func main() {
	history, err := getBashHistory()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(history)
}
