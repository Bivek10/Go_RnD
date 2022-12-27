package main

import (
	"bivek.com/fmt_api/operations"
)

func main() {
	runApp()
}
func runApp() {
	operations.Login("ram@gmail.com", "1234")
	//operations.Migrate()
}
