package util

import (
	"bufio"
	"log"
	"os"
)

type Parser func(text string)

func ReadLine(path string, parser Parser) {
	f, openErr := os.Open(path)
	if openErr != nil {
		log.Fatal(openErr)
	}
	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			log.Fatal(closeErr)
		}
	}()

	scanner := bufio.NewScanner(f)
	defer func() {
		scanErr := scanner.Err()
		if scanErr != nil {
			log.Fatal(scanErr)
		}
	}()
	for scanner.Scan() {
		parser(scanner.Text())
	}
}