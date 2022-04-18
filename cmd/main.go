package main

import "go-calendar-practice/go-calendar-practice/go-calendar/server"

func main() {
	const PORT = ":3000"

	server.Launch(PORT)
}
