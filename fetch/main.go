package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/pkg/errors"
)

func fetch(url string) (body []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "get failed")
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read failed")
	}
	return body, err
}

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "Please specify url")
		os.Exit(1)
	}

	for _, url := range os.Args[1:] {
		body, err := fetch(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s", body)
	}
}
