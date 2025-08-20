package tasksatu

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func OpenFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", errors.New("gagal membuka file")
	}
	defer file.Close()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	var result string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		panic("continue...")
	}

	return result, nil
}
