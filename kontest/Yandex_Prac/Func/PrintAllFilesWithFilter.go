package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// PrintAllFiles(".")
	PrintAllFilesWithFilter(".", "7-8kyu")
}

func PrintAllFilesWithFilter(path string, filter string) {
	// получаем список всех элементов в папке (и файлов, и директорий)
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("unable to get list of files", err)
		return
	}
	//  проходим по списку
	for _, f := range files {
		// получаем имя элемента
		// filepath.Join — функция, которая собирает путь к элементу с разделителями
		filename := filepath.Join(path, f.Name())
		// печатаем имя элемента
		if strings.Contains(filename, filter) {
			fmt.Println(filename)
		}
		if f.IsDir() {
			PrintAllFilesWithFilter(filename, filter)
		}

		// если элемент — директория, то вызываем для него рекурсивно ту же функцию

	}
}
