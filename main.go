package main

import (
	"tggcoin/cli"
	"tggcoin/db"
)

func main() {
	// defer 프로그램 종료시 실행
	defer db.Close()
	cli.Start()
}
