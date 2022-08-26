package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	paths := getFilePath("./")

	for _, file := range paths {
		fmt.Printf("[%s]\n", file)
		fp, err := os.Open(file)
		if err != nil {
			log.Fatalln(err)
		}
		defer fp.Close()

		// func NewScanner(r io.Reader) *Scanner
		scanner := bufio.NewScanner(fp)

		line_num := 1
		for scanner.Scan() {
			if strings.Contains(scanner.Text(), os.Args[1]) {
				fmt.Printf("%d行目: %s\n", line_num, scanner.Text())
			}
			line_num++
		}
		fmt.Println("----------------------")
	}
	/*
		str := "banana apple lemon"
		word := "apple"
		index := strings.Index(str, word)
		fmt.Println(index)
		fmt.Println(string(str[index : index+len(word)]))
	*/
}
