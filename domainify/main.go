package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"
)

const allowedChars = "abcdefghijklmnopqrstuvwxyz0123456789_-"

func main() {
	var tld = flag.String("tld", "com", "ドメイン")
	flag.Parse()

	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		text := strings.ToLower(s.Text())
		var newText []int32

		for _, r := range text {
			if unicode.IsSpace(r) {
				r = '-'
			}

			if !strings.ContainsRune(allowedChars, r) {
				continue
			}
			newText = append(newText, r)

		}

		fmt.Println(string(newText) + "." + *tld)
	}
}
