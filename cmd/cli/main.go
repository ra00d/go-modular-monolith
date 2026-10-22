/*
Copyright © 2026 Ra'ad A. Noman
*/
package main

import (
	"modular-monolith/internal/cli"
	"modular-monolith/internal/modules/auth"
)

func main() {
	cli.Execute(auth.Module)
}
