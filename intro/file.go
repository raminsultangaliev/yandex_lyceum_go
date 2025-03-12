package main 

import (
	"os"
	"bufio"
	"time"
	"errors"
)

func ReadContent(filename string) string {
	date, err := os.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(date)
}

func LineByNum(inputFilename string, lineNum int) string {
	f, err := os.Open(inputFilename)
	if err != nil {
		return ""
	}
	defer f.Close()

	fileScanner := bufio.NewScanner(f) 
	currentLine := 	0
	for fileScanner.Scan() { 
		if currentLine == lineNum {
			return fileScanner.Text()
		}
		currentLine++
	}
	return ""
}

func CopyFilePart(inputFilename, outFileName string, startPos int) error {
	inputFile, err := os.Open(inputFilename)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	inputFile.Seek(int64(startPos), 0)
	buffer := make([]byte, 1024)
	n, err := inputFile.Read(buffer)
	if err != nil {
		return err
	}

	outputFile, err := os.Create(outFileName)
	if err != nil {
			return err
	}
	defer outputFile.Close()

	_, err = outputFile.WriteString(string(buffer[:n])) 
	if err != nil {
		return err
	}

	return nil
}

func ModifyFile(filename string, pos int, val string) error {
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
 
	_, err = file.Seek(int64(pos), 0)
	if err != nil {
		return err
	}
 
	_, err = file.WriteString(val)
	if err != nil {
		return err
	}
 
	return nil
 }

 func ExtractLog(inputFileName string, start, end time.Time) ([]string, error) {
	file, err := os.Open(inputFileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var logs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		dateStr := line[:10]
		date, err := time.Parse("02.01.2006", dateStr)
		if err != nil {
			return nil, err
		}
		if date.After(start) && date.Before(end) || date.Equal(start) || date.Equal(end) {
			logs = append(logs, line)
		}
	}

	if len(logs) == 0 {
		return nil, errors.New("no logs in the specified range")
	}

	return logs, nil
}