package main

import (
	"fmt"

	"bivek.com/fmt_api/helper"
)

func main() {
	runApp()
}
func runApp() {
	helper.ConnectDB()
	fmt.Println("fmt running")
}
