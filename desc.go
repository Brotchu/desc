package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arguments := os.Args

	if arguments[1] == "init" {
		if _, err := os.Stat(".desc"); os.IsNotExist(err) {
			err = ioutil.WriteFile(".desc", []byte(""), 0644)
			check(err, "creating desc file")
			fmt.Println("Directory initiated..")
		} else {
			fmt.Println("Directory is already initiated..")
		}
	} else if arguments[1] == "-m" {
		if _, err := os.Stat(arguments[3]); os.IsNotExist(err) {
			fmt.Println("Cannot find file", arguments[3])
		} else {
			addMessage(arguments[2], arguments[3])
		}
	} else if arguments[1] == "-l" {
		listDescription()
	}
}

func check(err error, msg string) {
	if err != nil {
		fmt.Println("[Err ]: ", msg)
		fmt.Println(err)
		panic(err)
	}
}
