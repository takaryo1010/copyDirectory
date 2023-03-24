package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func check_copied_directory() error {
	_, err := os.Stat("copied_directory")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return nil
}

func make_directory_or_file(filename string) {
	var new_path string
	filepath.Walk("copied_directory", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		new_path = strings.Replace(path, "copied_directory", directory_name, -1)
		if strings.Index(info.Name(), ".") == -1 {
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
