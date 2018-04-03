package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/jinkwangchoi/codeinterview/tile"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: codeinterview inputfilename")
	}

	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
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
