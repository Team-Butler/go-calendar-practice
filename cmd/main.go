package main

import "go-calendar-practice/go-calendar-practice/go-calendar/server"

func main() {
	const PORT = ":8080"

	server.Launch(PORT)
}
