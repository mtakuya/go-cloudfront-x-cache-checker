package gocloudfrontxcachechecker

import (
	"fmt"
	"net/http"
	"os"
)

func Check() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "argument error")
		os.Exit(1)
	}

	url := os.Args[1]
	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	xCache := resp.Header.Get("x-cache")

	if xCache == "" {
		fmt.Fprintln(os.Stderr, "x-cache header not found")
	} else {
		fmt.Println(xCache)
	}
}
