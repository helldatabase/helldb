package server

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const filename = "keys.txt"

var Keys map[string]string

func GuardInit() {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	var text string
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		text = scanner.Text()
		r := csv.NewReader(strings.NewReader(text))
		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatalln(err)
			}
			Keys[record[0]] = record[1]
		}
	}
	_ = f.Close()
}

func WriteKey(username, password string) {
	Keys[username] = hash(password)
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalln(err)
	}
	for oldUsername, oldHashedPassword := range Keys {
		_, _ = f.WriteString(fmt.Sprintf("%s,%s\n", oldUsername, oldHashedPassword))
	}
}

func Verify(username, password string) bool {
	if val, ok := Keys[username]; ok {
		return hash(password) == val
	}
	return false
}
