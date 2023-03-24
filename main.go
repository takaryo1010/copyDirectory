package main

import (
	"log"
	"os"
)

var (
	directory_name string
)

func init() {
	if len(os.Args) == 2 {
		directory_name = os.Args[1]
	} else {
		log.Println("Add the name of the new directory you want to add like in the example.")
		log.Println("[example] go run . dir_name")
		os.Exit(1)
	}
}

func main() {
	check_copied_directory()
	make_directory_or_file(directory_name)

}
