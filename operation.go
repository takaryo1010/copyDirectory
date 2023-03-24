package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func check_original_directory() error {
	_, err := os.Stat("original_directory")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return nil
}

func make_directory_or_file(filename string) {
	var new_path string
	filepath.Walk("original_directory", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		new_path = strings.Replace(path, "original_directory", filename, -1)
		if info.IsDir() {
			err := os.MkdirAll(new_path, 0755)
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
		} else {
			dst, err := os.Create(new_path)
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
			src, err := os.Open(path)
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
			_, err = io.Copy(dst, src)
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
		}
		return nil
	})

}
