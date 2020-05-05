package pkg

import (
	"github.com/nymo/gonsole/internal/command"
)

type Console interface {
	RegisterCommand(command command.Command)
	Execute(arguments string) (int, error)
}

type console struct {
	commandList   command.List
	commandFinder command.Finder
}

func (console *console) RegisterCommand(command command.Command) {
	console.commandList.Add(command)
}

func (console *console) Execute(arguments string) (int, error) {
	com, err := console.commandFinder.FindCommandToExecute(arguments)
	if err != nil {
		return 0, err
	}
	return com.Print(), nil
}

func NewConsole() Console {
	list := command.NewList()
	finder := command.NewFinder(list)

	return &console{list, finder}
}
