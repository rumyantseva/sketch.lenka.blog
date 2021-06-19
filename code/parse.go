package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/rumyantseva/sketch.lenka.blog/watercolor"
)

func main() {
	log.SetOutput(os.Stdout)

	defer func() {
		recover()
	}()

	file, err := os.Open("./paints.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Print(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	count := 1

	paintsNr := 250
	paints := make([]watercolor.Paint, paintsNr)
	for i := 0; i < paintsNr; i++ {
		paint := watercolor.Paint{}

		// Scan title and ID
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		log.Println(line)
		parts := strings.Split(line, "№")
		if len(parts) != 2 {
			log.Panicf("Can't parse № from the line `%s`", line)
		}
		count++

		// Scan paint properties
		if !scanner.Scan() {
			break
		}
		line = scanner.Text()
		err := paint.ParseWhiteNights(line)
		if err != nil {
			log.Panicf("Can't parse properties: `%v` from the string `%s`", err, line)
		}
		count++

		paints[i] = paint

		// The third line is empty
		if !scanner.Scan() {
			break
		}
		count++
	}

	if err := scanner.Err(); err != nil {
		log.Panic(err)
	}

	log.Printf("Parsed %d lines\n", count)

	fr, _ := os.Create("./result.json")
	json.NewEncoder(fr).Encode(paints)
	fr.Close()
}
