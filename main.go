package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	fileName = flag.String("f", "", "file name")
)

func ccs(s string) (int, bool) {
	for i := 1; i < 32; i++ {
		var cc = []byte{byte(i)}
		if strings.Contains(s, string(cc)) {
			return i, true
		}
	}

	if strings.Contains(s, string(0x7f)) {
		return 0x7f, true
	}
	return -1, false
}

func main() {
	flag.Parse()
	f, err := os.Open(*fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	ln := 0
	for scanner.Scan() {
		if cCode, result := ccs(scanner.Text()); result {
			fmt.Printf("line: %d, char code: %d\n", ln+1, cCode)
		}
		ln++
	}
}
