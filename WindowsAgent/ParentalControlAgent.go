package main

import (
	"github.com/ideastorm/w32"
)

func logout() {
	w32.ExitWindowsEx(0,0)
}

func getUserName() string {
	var userName string
	userName = w32.GetUserNameEx
	return "guest"
}

func main() {
}
