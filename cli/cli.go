package cli

import (
	"flag"
	"fmt"
	"nomadcoin/explorer"
	"nomadcoin/rest"
	"os"
)

func usage() {
	fmt.Printf("Welcome to TGG coin\n\n")
	fmt.Printf("Please use the following flags: \n\n")
	fmt.Printf("-port:      Set the PORT of the server\n")
	fmt.Printf("-mode:      Choose between 'html' and 'rest\n")
	os.Exit(0)
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}
}
