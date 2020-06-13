package main

import (
	"bufio"
	"context"
	"log"
	"os"

	"github.com/eskriett/eml2img"
)

const numRequiredArgs = 3

func main() {
	if len(os.Args) < numRequiredArgs {
		log.Fatal("eml2img <INPUT> <OUTPUT>")
	}

	e2i, err := eml2img.New()
	check(err)

	input, err := os.Open(os.Args[1])
	check(err)

	defer func() { check(input.Close()) }()

	output, err := os.Create(os.Args[2])
	check(err)

	defer func() { check(output.Close()) }()

	w := bufio.NewWriter(output)

	check(e2i.Convert(context.Background(), input, w))
	check(w.Flush())
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
