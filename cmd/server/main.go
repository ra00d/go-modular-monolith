package main

import (
	"modular-monolith/internal/auth"
	"modular-monolith/lib/app"
)

func main() {
	app := app.NewApp(auth.Module)
	app.Run()
}
