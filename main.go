package main

import (
	"tggcoin/cli"
	"tggcoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
