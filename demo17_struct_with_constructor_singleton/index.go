package main

import (
	"demo17_struct_with_constructor/employee"
)

func main() {
	e := employee.Init(
		"Sam",
		"Adolf",
		30,
		20)

	e.LeavesRemaining()

	e = employee.Init(
		"Kan",
		"Adolf",
		30,
		20)

	e.LeavesRemaining()
}

//run command go mod init demo15 for create .mod
