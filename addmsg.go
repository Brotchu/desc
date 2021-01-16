package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func addMessage(message string, filename string) {
	fmt.Println("File :", filename, message)
	if _, err := os.Stat(".desc"); os.IsNotExist(err) {
		fmt.Println("Directory not initiated.")

	} else {
		//Open file to read
		file, err := os.OpenFile(".desc", os.O_RDWR, 0644)
		check(err, "Opening file")
		defer file.Close()

		var lines []string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		fileFlag := false //to check if entry exists
		for index, line := range lines {
			f_name := strings.Split(line, ":")[0]
			if f_name == filename {
				lines[index] = f_name + ":" + message
				fileFlag = true
			}
		}
		//If new entry
		if fileFlag == false {
			lines = append(lines, filename+":"+message)
		}

		// fmt.Println("Description", lines)

		file.Truncate(0)
		for _, line := range lines {
			file.WriteString(line + "\n")
		}
	}
}
