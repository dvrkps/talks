// +build OMIT

package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("https://golang.org")
	defer res.Body.Close() // HL
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, res.Body)
}