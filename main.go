package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

const url = "http://s3.amazonaws.com/sem3_misc/CategoryTree_Semantics3.txt"

var rxLine = regexp.MustCompile(` *(\d)\. *(.+) \((\d+)\).*`)

func requestCatLines(url string) ([]byte, error) {
	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("requestCatLines: no 200 response; got %v", resp.StatusCode)
	}

	bodyBs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return bodyBs, nil
}

func parseAndPrint(bs []byte) error {
	var cat1, cat2, cat3, cat4 string
	scanner := bufio.NewScanner(bytes.NewReader(bs))
	for scanner.Scan() {
		results := rxLine.FindAllStringSubmatch(scanner.Text(), -1)
		for _, result := range results {
			num, cat, id := result[1], result[2], result[3]
			switch num {
			case "1":
				cat1 = cat
				cat2, cat3, cat4 = "", "", ""
			case "2":
				cat2 = cat
				cat3, cat4 = "", ""
			case "3":
				cat3 = cat
				cat4 = ""
			case "4":
				cat4 = cat
			}
			fmt.Printf("%v\t%v\t%v\t%v\t%v\n", id, cat1, cat2, cat3, cat4)
			break
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func main() {
	bs, err := requestCatLines(url)
	if err != nil {
		log.Fatal(err)
	}
	if err := parseAndPrint(bs); err != nil {
		log.Fatal(err)
	}
}
