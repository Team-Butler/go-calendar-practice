package main

import "go-calendar-practice/go-calendar-practice/go-calendar/services"

func main() {
	const PORT = ":3000"

	services.LaunchServer(PORT)
}
