package main

import (
	"bufio"
	"fmt"
	"github.com/jinkwangchoi/codeinterview/tile"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: codeinterview inputfilename")
		os.Exit(1)
	}

	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	var tiles []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tiles = append(tiles, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	g := tile.NewGroup(tiles)
	fmt.Print(g.Password())
}
