package main

import "github.com/jtaczanowski/tcp-pinger/cmd/tcp-pinger/factory"

func main() {
	app := factory.App{}
	app.Initialize()
	app.Run()
}
