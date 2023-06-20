package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

func main() {
	tok := os.Getenv("ACCESSTOKEN")
	if tok == "" {
		bin, err := os.Executable()
		if err != nil {
			bin = "jwtdec"
		}
		fmt.Fprintln(os.Stderr, "Please run this tool with the to-be-decoded access token set as the ACCESSTOKEN env variable, e.g.:")
		fmt.Fprintf(os.Stderr, "$ ACCESSTOKEN=\"...\" %s\n", bin)
		os.Exit(-1)
	}
	subtoks := strings.Split(tok, ".")
	if len(subtoks) != 3 {
		fmt.Fprintf(os.Stderr, "wrong number of subtokens: expected 3, got %d\n", len(subtoks))
	}
	header := subtoks[0]
	payload := subtoks[1]
	hd, err := base64.StdEncoding.DecodeString(header)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not base64 decode header: %s\n", err)
	} else {
		fmt.Println(string(hd))
	}
	pd, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not base64 decode payload: %s\n", err)
	} else {
		fmt.Println(string(pd))
	}
}
