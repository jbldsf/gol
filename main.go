package main

import "gol/server"

func main() {
	go server.HTTPS()
	server.HTTP()
}
