package main

import (
	"demo16_struct_with_constructor/employee"
)

func main() {
	e := employee.New(
		"Sam",
		"Adolf",
		30,
		20)

	e.LeavesRemaining()
}

//run command go mod init demo15 for create .mod
