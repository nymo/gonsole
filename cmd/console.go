package main

import (
	"github.com/nymo/gonsole/examples"
	"github.com/nymo/gonsole/pkg"
	"os"
)

func main() {
	console := pkg.NewConsole()
	console.RegisterCommand(examples.Foo{})
	console.RegisterCommand(examples.Bar{})
	returnCode, err := console.Execute("cache:clear") //bekommt an dieser stelle dann die arguments übergeben
	if err != nil {
		print(err.Error())
	}

	os.Exit(returnCode)
}
