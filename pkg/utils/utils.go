package utils

import (
"bufio"
"os"
"strings"
)

func ReadToken(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	token := strings.TrimSpace(scanner.Text())
	
	if err := scanner.Err(); err != nil {
		return "", err
	}
	
	return token, nil
}
