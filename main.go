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

func ccs(s string) bool {
	for i := 1; i < 32; i++ {
		var cc = []byte{byte(i)}
		if strings.Contains(s, string(cc)) {
			return true
		}
	}

	return strings.Contains(s, string(0x7f))
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
		if ccs(scanner.Text()) {
			fmt.Println(ln + 1)
		}
		ln++
	}
}
