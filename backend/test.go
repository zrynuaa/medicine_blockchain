package main

import (
	"github.com/scottocs/medicine_blockchain/backend/server"
	"github.com/scottocs/medicine_blockchain/backend/based"
)

func main()  {
	go based.StartTimer()
	server.Run()
}