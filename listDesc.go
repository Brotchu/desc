package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func listDescription() {
	if _, err := os.Stat(".desc"); os.IsNotExist(err) {
		fmt.Println("Directory not initiated.")
	} else {

		descMap := make(map[string]string)

		//TODO: make function to read file,
		file, err := os.OpenFile(".desc", os.O_RDWR, 0644)
		check(err, "Opening file")
		defer file.Close()

		var lines []string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		for _, line := range lines {
			details := strings.Split(line, ":")
			descMap[details[0]] = details[1]
		}

		cmd := exec.Command("ls", "-lrt")
		output, err := cmd.Output()
		check(err, "running ls command")

		for _, outline := range strings.Split(string(output), "\n")[1:] {
			tempLine := strings.Split(outline, " ")
			fmt.Println(outline + "\t ---- " + descMap[tempLine[len(tempLine)-1]])
		}
	}
}
